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

func configGinObjectsRouter(router gin.IRoutes) {
	router.GET("/objects", ConverHttpRouterToGin(GetAllObjects))
	router.POST("/objects", ConverHttpRouterToGin(AddObjects))
	router.GET("/objects/:argID", ConverHttpRouterToGin(GetObjects))
	router.PUT("/objects/:argID", ConverHttpRouterToGin(UpdateObjects))
	router.DELETE("/objects/:argID", ConverHttpRouterToGin(DeleteObjects))
}

// GetAllObjects is a function to get a slice of record(s) from objects table in the estuary database
// @Summary Get list of Objects
// @Tags Objects
// @Description GetAllObjects is a handler to get a slice of record(s) from objects table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Object}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /objects [get]
// http "http://localhost:3030/objects?page=0&pagesize=20" X-Api-User:user123
func GetAllObjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "objects", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllObjects(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetObjects is a function to get a single record from the objects table in the estuary database
// @Summary Get record from table Objects by  argID
// @Tags Objects
// 
// @Description GetObjects is a function to get a single record from the objects table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Object
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /objects/{argID} [get]
// http "http://localhost:3030/objects/1" X-Api-User:user123
func GetObjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "objects", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetObjects(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddObjects add to add a single record to objects table in the estuary database
// @Summary Add an record to objects table
// @Description add to add a single record to objects table in the estuary database
// @Tags Objects
// @Accept  json
// @Produce  json
// @Param Objects body model.Object true "Add Objects"
// @Success 200 {object} model.Object
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /objects [post]
// echo '{"id": 31,"cid": "FbbIkboFwFjwjgVAgqtDfgRLb","size": 67,"reads": 45,"last_access": "2088-10-29T18:32:13.860989039-04:00"}' | http POST "http://localhost:3030/objects" X-Api-User:user123
func AddObjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	objects := &model.Object{}

	if err := readJSON(r, objects); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := objects.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	objects.Prepare()

	if err := objects.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "objects", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	objects, _, err = dao.AddObjects(ctx, objects)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, objects)
}

// UpdateObjects Update a single record from objects table in the estuary database
// @Summary Update an record in table objects
// @Description Update a single record from objects table in the estuary database
// @Tags Objects
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Objects body model.Object true "Update Objects record"
// @Success 200 {object} model.Object
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /objects/{argID} [put]
// echo '{"id": 31,"cid": "FbbIkboFwFjwjgVAgqtDfgRLb","size": 67,"reads": 45,"last_access": "2088-10-29T18:32:13.860989039-04:00"}' | http PUT "http://localhost:3030/objects/1"  X-Api-User:user123
func UpdateObjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	objects := &model.Object{}
	if err := readJSON(r, objects); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := objects.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	objects.Prepare()

	if err := objects.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "objects", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	objects, _, err = dao.UpdateObjects(ctx,
		argID,
		objects)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, objects)
}

// DeleteObjects Delete a single record from objects table in the estuary database
// @Summary Delete a record from objects
// @Description Delete a single record from objects table in the estuary database
// @Tags Objects
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Object
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /objects/{argID} [delete]
// http DELETE "http://localhost:3030/objects/1" X-Api-User:user123
func DeleteObjects(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "objects", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteObjects(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
