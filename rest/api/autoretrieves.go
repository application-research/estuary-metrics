package api

import (
	"net/http"

	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configGinAutoretrievesRouter(router gin.IRoutes) {
	router.GET("/autoretrieves", ConverHttpRouterToGin(GetAllAutoretrieves))
	router.POST("/autoretrieves", ConverHttpRouterToGin(AddAutoretrieves))
	router.GET("/autoretrieves/:argID", ConverHttpRouterToGin(GetAutoretrieves))
	router.PUT("/autoretrieves/:argID", ConverHttpRouterToGin(UpdateAutoretrieves))
	router.DELETE("/autoretrieves/:argID", ConverHttpRouterToGin(DeleteAutoretrieves))
}

// GetAllAutoretrieves is a function to get a slice of record(s) from autoretrieves table in the estuary database
// @Summary Get list of Autoretrieves
// @Tags Autoretrieves
// @Description GetAllAutoretrieves is a handler to get a slice of record(s) from autoretrieves table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Autoretrieve}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /autoretrieves [get]
// http "http://localhost:3030/autoretrieves?page=0&pagesize=20" X-Api-User:user123
func GetAllAutoretrieves(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	page, err := readInt(r, "page", 0)
	if err != nil || page < 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := ValidateRequest(ctx, r, "autoretrieves", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAutoretrieves(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAutoretrieves is a function to get a single record from the autoretrieves table in the estuary database
// @Summary Get record from table Autoretrieves by  argID
// @Tags Autoretrieves
//
// @Description GetAutoretrieves is a function to get a single record from the autoretrieves table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Autoretrieve
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /autoretrieves/{argID} [get]
// http "http://localhost:3030/autoretrieves/1" X-Api-User:user123
func GetAutoretrieves(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "autoretrieves", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAutoretrieves(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAutoretrieves add to add a single record to autoretrieves table in the estuary database
// @Summary Add an record to autoretrieves table
// @Description add to add a single record to autoretrieves table in the estuary database
// @Tags Autoretrieves
// @Accept  json
// @Produce  json
// @Param Autoretrieves body model.Autoretrieve true "Add Autoretrieves"
// @Success 200 {object} model.Autoretrieve
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /autoretrieves [post]
// echo '{"id": 78,"created_at": "2189-06-27T20:49:09.654057738-04:00","updated_at": "2027-07-03T12:05:44.436177866-04:00","deleted_at": "2313-08-27T19:00:22.797210455-04:00","handle": "sqdtjZASNduRPMrExTEaBmxXH","token": "CNrJLhpYdHxwkHnccJxPmbGFj","last_connection": "2287-06-23T04:27:03.570478714-04:00","peer_id": "RbNrTOgFeoExgXcpZClEekogt","addresses": "JDrAETmltQNBfwKZvBMLfBPdm"}' | http POST "http://localhost:3030/autoretrieves" X-Api-User:user123
func AddAutoretrieves(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	autoretrieves := &model.Autoretrieve{}

	if err := readJSON(r, autoretrieves); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := autoretrieves.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	autoretrieves.Prepare()

	if err := autoretrieves.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "autoretrieves", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	autoretrieves, _, err = dao.AddAutoretrieves(ctx, autoretrieves)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, autoretrieves)
}

// UpdateAutoretrieves Update a single record from autoretrieves table in the estuary database
// @Summary Update an record in table autoretrieves
// @Description Update a single record from autoretrieves table in the estuary database
// @Tags Autoretrieves
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Autoretrieves body model.Autoretrieve true "Update Autoretrieves record"
// @Success 200 {object} model.Autoretrieve
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /autoretrieves/{argID} [put]
// echo '{"id": 78,"created_at": "2189-06-27T20:49:09.654057738-04:00","updated_at": "2027-07-03T12:05:44.436177866-04:00","deleted_at": "2313-08-27T19:00:22.797210455-04:00","handle": "sqdtjZASNduRPMrExTEaBmxXH","token": "CNrJLhpYdHxwkHnccJxPmbGFj","last_connection": "2287-06-23T04:27:03.570478714-04:00","peer_id": "RbNrTOgFeoExgXcpZClEekogt","addresses": "JDrAETmltQNBfwKZvBMLfBPdm"}' | http PUT "http://localhost:3030/autoretrieves/1"  X-Api-User:user123
func UpdateAutoretrieves(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	autoretrieves := &model.Autoretrieve{}
	if err := readJSON(r, autoretrieves); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := autoretrieves.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	autoretrieves.Prepare()

	if err := autoretrieves.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "autoretrieves", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	autoretrieves, _, err = dao.UpdateAutoretrieves(ctx,
		argID,
		autoretrieves)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, autoretrieves)
}

// DeleteAutoretrieves Delete a single record from autoretrieves table in the estuary database
// @Summary Delete a record from autoretrieves
// @Description Delete a single record from autoretrieves table in the estuary database
// @Tags Autoretrieves
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Autoretrieve
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /autoretrieves/{argID} [delete]
// http DELETE "http://localhost:3030/autoretrieves/1" X-Api-User:user123
func DeleteAutoretrieves(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "autoretrieves", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAutoretrieves(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
