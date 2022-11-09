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

func configGinShuttlesRouter(router gin.IRoutes) {
	router.GET("/shuttles", ConverHttpRouterToGin(GetAllShuttles))
	router.GET("/shuttles/count", ConverHttpRouterToGin(GetNumberOfRegisteredShuttles))
	router.GET("/shuttles/:argID", ConverHttpRouterToGin(GetShuttles))
}

func GetNumberOfRegisteredShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "shuttles", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var totalRows int64
	dao.DB.Model(&model.Shuttle{}).Count(&totalRows)

	writeJSON(ctx, w, totalRows)
}

// GetAllShuttles is a function to get a slice of record(s) from shuttles table in the estuary database
// @Summary Get list of Shuttles
// @Tags Shuttles
// @Description GetAllShuttles is a handler to get a slice of record(s) from shuttles table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Shuttle}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shuttles [get]
// http "http://localhost:3030/shuttles?page=0&pagesize=20" X-Api-User:user123
func GetAllShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "shuttles", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllShuttles(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetShuttles is a function to get a single record from the shuttles table in the estuary database
// @Summary Get record from table Shuttles by  argID
// @Tags Shuttles
//
// @Description GetShuttles is a function to get a single record from the shuttles table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Shuttle
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /shuttles/{argID} [get]
// http "http://localhost:3030/shuttles/1" X-Api-User:user123
func GetShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shuttles", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetShuttles(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddShuttles add to add a single record to shuttles table in the estuary database
// @Summary Add an record to shuttles table
// @Description add to add a single record to shuttles table in the estuary database
// @Tags Shuttles
// @Accept  json
// @Produce  json
// @Param Shuttles body model.Shuttle true "Add Shuttles"
// @Success 200 {object} model.Shuttle
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shuttles [post]
// echo '{"id": 22,"created_at": "2145-07-22T19:17:01.371381383-04:00","updated_at": "2266-09-09T22:16:24.664940564-04:00","deleted_at": "2082-11-05T17:51:01.923398803-05:00","handle": "SNYtRXPLWPJlLMvIuJpdcnpiT","token": "krTOakeSGgTBYJfHgNtBmIjrH","last_connection": "2276-08-21T16:55:00.358752098-04:00","host": "pwVbwWCKsKfTbpMZHfocyuiKr","peer_id": "xoffxJqXhXSQtYnTQOTWFkEvw","open": true,"private": true,"priority": 53}' | http POST "http://localhost:3030/shuttles" X-Api-User:user123
func AddShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	shuttles := &model.Shuttle{}

	if err := readJSON(r, shuttles); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shuttles.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shuttles.Prepare()

	if err := shuttles.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shuttles", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	shuttles, _, err = dao.AddShuttles(ctx, shuttles)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shuttles)
}

// UpdateShuttles Update a single record from shuttles table in the estuary database
// @Summary Update an record in table shuttles
// @Description Update a single record from shuttles table in the estuary database
// @Tags Shuttles
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Shuttles body model.Shuttle true "Update Shuttles record"
// @Success 200 {object} model.Shuttle
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /shuttles/{argID} [put]
// echo '{"id": 22,"created_at": "2145-07-22T19:17:01.371381383-04:00","updated_at": "2266-09-09T22:16:24.664940564-04:00","deleted_at": "2082-11-05T17:51:01.923398803-05:00","handle": "SNYtRXPLWPJlLMvIuJpdcnpiT","token": "krTOakeSGgTBYJfHgNtBmIjrH","last_connection": "2276-08-21T16:55:00.358752098-04:00","host": "pwVbwWCKsKfTbpMZHfocyuiKr","peer_id": "xoffxJqXhXSQtYnTQOTWFkEvw","open": true,"private": true,"priority": 53}' | http PUT "http://localhost:3030/shuttles/1"  X-Api-User:user123
func UpdateShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shuttles := &model.Shuttle{}
	if err := readJSON(r, shuttles); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := shuttles.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	shuttles.Prepare()

	if err := shuttles.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "shuttles", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	shuttles, _, err = dao.UpdateShuttles(ctx,
		argID,
		shuttles)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, shuttles)
}

// DeleteShuttles Delete a single record from shuttles table in the estuary database
// @Summary Delete a record from shuttles
// @Description Delete a single record from shuttles table in the estuary database
// @Tags Shuttles
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Shuttle
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /shuttles/{argID} [delete]
// http DELETE "http://localhost:3030/shuttles/1" X-Api-User:user123
func DeleteShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "shuttles", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteShuttles(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
