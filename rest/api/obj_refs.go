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

func configGinObjRefsRouter(router gin.IRoutes) {
	router.GET("/objrefs", ConverHttpRouterToGin(GetAllObjRefs))
	router.POST("/objrefs", ConverHttpRouterToGin(AddObjRefs))
	router.GET("/objrefs/:argID", ConverHttpRouterToGin(GetObjRefs))
	router.PUT("/objrefs/:argID", ConverHttpRouterToGin(UpdateObjRefs))
	router.DELETE("/objrefs/:argID", ConverHttpRouterToGin(DeleteObjRefs))
}

// GetAllObjRefs is a function to get a slice of record(s) from obj_refs table in the estuary database
// @Summary Get list of ObjRefs
// @Tags ObjRefs
// @Description GetAllObjRefs is a handler to get a slice of record(s) from obj_refs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ObjRef}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /objrefs [get]
// http "http://localhost:3030/objrefs?page=0&pagesize=20" X-Api-User:user123
func GetAllObjRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "obj_refs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllObjRefs(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetObjRefs is a function to get a single record from the obj_refs table in the estuary database
// @Summary Get record from table ObjRefs by  argID
// @Tags ObjRefs
//
// @Description GetObjRefs is a function to get a single record from the obj_refs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ObjRef
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /objrefs/{argID} [get]
// http "http://localhost:3030/objrefs/1" X-Api-User:user123
func GetObjRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "obj_refs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetObjRefs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddObjRefs add to add a single record to obj_refs table in the estuary database
// @Summary Add an record to obj_refs table
// @Description add to add a single record to obj_refs table in the estuary database
// @Tags ObjRefs
// @Accept  json
// @Produce  json
// @Param ObjRefs body model.ObjRef true "Add ObjRefs"
// @Success 200 {object} model.ObjRef
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /objrefs [post]
// echo '{"id": 83,"content": 35,"object": 99,"offloaded": 88}' | http POST "http://localhost:3030/objrefs" X-Api-User:user123
func AddObjRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	objrefs := &model.ObjRef{}

	if err := readJSON(r, objrefs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := objrefs.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	objrefs.Prepare()

	if err := objrefs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "obj_refs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	objrefs, _, err = dao.AddObjRefs(ctx, objrefs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, objrefs)
}

// UpdateObjRefs Update a single record from obj_refs table in the estuary database
// @Summary Update an record in table obj_refs
// @Description Update a single record from obj_refs table in the estuary database
// @Tags ObjRefs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ObjRefs body model.ObjRef true "Update ObjRefs record"
// @Success 200 {object} model.ObjRef
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /objrefs/{argID} [put]
// echo '{"id": 83,"content": 35,"object": 99,"offloaded": 88}' | http PUT "http://localhost:3030/objrefs/1"  X-Api-User:user123
func UpdateObjRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	objrefs := &model.ObjRef{}
	if err := readJSON(r, objrefs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := objrefs.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	objrefs.Prepare()

	if err := objrefs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "obj_refs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	objrefs, _, err = dao.UpdateObjRefs(ctx,
		argID,
		objrefs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, objrefs)
}

// DeleteObjRefs Delete a single record from obj_refs table in the estuary database
// @Summary Delete a record from obj_refs
// @Description Delete a single record from obj_refs table in the estuary database
// @Tags ObjRefs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ObjRef
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /objrefs/{argID} [delete]
// http DELETE "http://localhost:3030/objrefs/1" X-Api-User:user123
func DeleteObjRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "obj_refs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteObjRefs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
