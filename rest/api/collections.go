package api

import (
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configGinCollectionsRouter(router gin.IRoutes) {
	router.GET("/collections", ConverHttpRouterToGin(GetAllCollections))
	router.POST("/collections", ConverHttpRouterToGin(AddCollections))
	router.GET("/collections/:argID", ConverHttpRouterToGin(GetCollections))
	router.PUT("/collections/:argID", ConverHttpRouterToGin(UpdateCollections))
	router.DELETE("/collections/:argID", ConverHttpRouterToGin(DeleteCollections))
}

// GetAllCollections is a function to get a slice of record(s) from collections table in the estuary database
// @Summary Get list of Collections
// @Tags Collections
// @Description GetAllCollections is a handler to get a slice of record(s) from collections table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Collection}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /collections [get]
// http "http://localhost:3030/collections?page=0&pagesize=20" X-Api-User:user123
func GetAllCollections(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "collections", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCollections(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetCollections is a function to get a single record from the collections table in the estuary database
// @Summary Get record from table Collections by  argID
// @Tags Collections
//
// @Description GetCollections is a function to get a single record from the collections table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Collection
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /collections/{argID} [get]
// http "http://localhost:3030/collections/1" X-Api-User:user123
func GetCollections(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "collections", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCollections(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddCollections add to add a single record to collections table in the estuary database
// @Summary Add an record to collections table
// @Description add to add a single record to collections table in the estuary database
// @Tags Collections
// @Accept  json
// @Produce  json
// @Param Collections body model.Collection true "Add Collections"
// @Success 200 {object} model.Collection
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /collections [post]
// echo '{"id": 7,"created_at": "2125-03-11T07:32:05.249809045-04:00","uuid": "yyZEgJrLRRBdDCgMslbEAlDWy","name": "mENXrtPYUsoUrCtKSUMsiJHrm","description": "mjlNQrPMrhokGacKduAAeNyyb","user_id": 67,"c_id": "shLJtDZHdQvInOmYXZbSDiUOQ"}' | http POST "http://localhost:3030/collections" X-Api-User:user123
func AddCollections(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	collections := &model.Collection{}

	if err := readJSON(r, collections); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := collections.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	collections.Prepare()

	if err := collections.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "collections", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	collections, _, err = dao.AddCollections(ctx, collections)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, collections)
}

// UpdateCollections Update a single record from collections table in the estuary database
// @Summary Update an record in table collections
// @Description Update a single record from collections table in the estuary database
// @Tags Collections
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Collections body model.Collection true "Update Collections record"
// @Success 200 {object} model.Collection
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /collections/{argID} [put]
// echo '{"id": 7,"created_at": "2125-03-11T07:32:05.249809045-04:00","uuid": "yyZEgJrLRRBdDCgMslbEAlDWy","name": "mENXrtPYUsoUrCtKSUMsiJHrm","description": "mjlNQrPMrhokGacKduAAeNyyb","user_id": 67,"c_id": "shLJtDZHdQvInOmYXZbSDiUOQ"}' | http PUT "http://localhost:3030/collections/1"  X-Api-User:user123
func UpdateCollections(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	collections := &model.Collection{}
	if err := readJSON(r, collections); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := collections.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	collections.Prepare()

	if err := collections.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "collections", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	collections, _, err = dao.UpdateCollections(ctx,
		argID,
		collections)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, collections)
}

// DeleteCollections Delete a single record from collections table in the estuary database
// @Summary Delete a record from collections
// @Description Delete a single record from collections table in the estuary database
// @Tags Collections
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Collection
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /collections/{argID} [delete]
// http DELETE "http://localhost:3030/collections/1" X-Api-User:user123
func DeleteCollections(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "collections", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCollections(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
