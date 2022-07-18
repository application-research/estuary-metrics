package api

import (
	"net/http"
	"time"

	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configGinContentsRouter(router gin.IRoutes) {
	router.GET("/contents", ConverHttpRouterToGin(GetAllContents))
	router.POST("/contents", ConverHttpRouterToGin(AddContents))
	router.GET("/contents/:argID", ConverHttpRouterToGin(GetContents))
	router.PUT("/contents/:argID", ConverHttpRouterToGin(UpdateContents))
	router.DELETE("/contents/:argID", ConverHttpRouterToGin(DeleteContents))
}

// GetAllContents is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description GetAllContents is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Content}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contents [get]
// http "http://localhost:3030/contents?page=0&pagesize=20" X-Api-User:user123
func GetAllContents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "contents", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContents(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetContents is a function to get a single record from the contents table in the estuary database
// @Summary Get record from table Contents by  argID
// @Tags Contents
//
// @Description GetContents is a function to get a single record from the contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Content
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contents/{argID} [get]
// http "http://localhost:3030/contents/1" X-Api-User:user123
func GetContents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "contents", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.Cacher.Get("getContents", time.Minute*2, func() (interface{}, error) {
		return dao.GetContents(ctx, argID)
	})

	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContents add to add a single record to contents table in the estuary database
// @Summary Add an record to contents table
// @Description add to add a single record to contents table in the estuary database
// @Tags Contents
// @Accept  json
// @Produce  json
// @Param Contents body model.Content true "Add Contents"
// @Success 200 {object} model.Content
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contents [post]
// echo '{"id": 93,"created_at": "2057-12-04T18:53:04.559551978-05:00","updated_at": "2071-01-09T11:00:13.333406634-05:00","deleted_at": "2289-08-11T03:08:57.420273964-04:00","cid": "KXbBuOgxMWsxQkLgxiGFBLrHu","name": "xhdhFyHbMxCSxndMHwRUxVgRU","user": "ZTQPqArWKUMaAVtXZVKBbJCgg","size": 98,"active": false,"description": "KcQTwjHtntnWxyWGZiaIWGDto","offloaded": true,"user_id": 99,"replication": 44,"aggregated_in": 35,"aggregate": false,"pinning": true,"pin_meta": "iDIhMqhTEvVAsrMeQxGpLIbqF","location": "xjrdXRYNwmmNdvtufUpNOGgbV","failed": true,"dag_split": true,"split_from": 68,"type": 9,"replace": false,"origins": "LvoLfwmEZraeTUdSARYNSPNjN"}' | http POST "http://localhost:3030/contents" X-Api-User:user123
func AddContents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contents := &model.Content{}

	if err := readJSON(r, contents); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contents.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contents.Prepare()

	if err := contents.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "contents", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contents, _, err = dao.AddContents(ctx, contents)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contents)
}

// UpdateContents Update a single record from contents table in the estuary database
// @Summary Update an record in table contents
// @Description Update a single record from contents table in the estuary database
// @Tags Contents
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Contents body model.Content true "Update Contents record"
// @Success 200 {object} model.Content
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contents/{argID} [put]
// echo '{"id": 93,"created_at": "2057-12-04T18:53:04.559551978-05:00","updated_at": "2071-01-09T11:00:13.333406634-05:00","deleted_at": "2289-08-11T03:08:57.420273964-04:00","cid": "KXbBuOgxMWsxQkLgxiGFBLrHu","name": "xhdhFyHbMxCSxndMHwRUxVgRU","user": "ZTQPqArWKUMaAVtXZVKBbJCgg","size": 98,"active": false,"description": "KcQTwjHtntnWxyWGZiaIWGDto","offloaded": true,"user_id": 99,"replication": 44,"aggregated_in": 35,"aggregate": false,"pinning": true,"pin_meta": "iDIhMqhTEvVAsrMeQxGpLIbqF","location": "xjrdXRYNwmmNdvtufUpNOGgbV","failed": true,"dag_split": true,"split_from": 68,"type": 9,"replace": false,"origins": "LvoLfwmEZraeTUdSARYNSPNjN"}' | http PUT "http://localhost:3030/contents/1"  X-Api-User:user123
func UpdateContents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contents := &model.Content{}
	if err := readJSON(r, contents); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contents.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contents.Prepare()

	if err := contents.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "contents", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contents, _, err = dao.UpdateContents(ctx,
		argID,
		contents)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contents)
}

// DeleteContents Delete a single record from contents table in the estuary database
// @Summary Delete a record from contents
// @Description Delete a single record from contents table in the estuary database
// @Tags Contents
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Content
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contents/{argID} [delete]
// http DELETE "http://localhost:3030/contents/1" X-Api-User:user123
func DeleteContents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "contents", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContents(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
