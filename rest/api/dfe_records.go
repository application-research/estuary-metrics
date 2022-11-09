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

func configGinDfeRecordsRouter(router gin.IRoutes) {
	router.GET("/dferecords", ConverHttpRouterToGin(GetAllDfeRecords))
	router.GET("/dferecords/:argID", ConverHttpRouterToGin(GetDfeRecords))
}

// GetAllDfeRecords is a function to get a slice of record(s) from dfe_records table in the estuary database
// @Summary Get list of DfeRecords
// @Tags DfeRecords
// @Description GetAllDfeRecords is a handler to get a slice of record(s) from dfe_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.DfeRecord}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dferecords [get]
// http "http://localhost:3030/dferecords?page=0&pagesize=20" X-Api-User:user123
func GetAllDfeRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "dfe_records", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDfeRecords(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetDfeRecords is a function to get a single record from the dfe_records table in the estuary database
// @Summary Get record from table DfeRecords by  argID
// @Tags DfeRecords
//
// @Description GetDfeRecords is a function to get a single record from the dfe_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.DfeRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /dferecords/{argID} [get]
// http "http://localhost:3030/dferecords/1" X-Api-User:user123
func GetDfeRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "dfe_records", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDfeRecords(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddDfeRecords add to add a single record to dfe_records table in the estuary database
// @Summary Add an record to dfe_records table
// @Description add to add a single record to dfe_records table in the estuary database
// @Tags DfeRecords
// @Accept  json
// @Produce  json
// @Param DfeRecords body model.DfeRecord true "Add DfeRecords"
// @Success 200 {object} model.DfeRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dferecords [post]
// echo '{"id": 27,"created_at": "2180-11-16T17:21:14.605962774-05:00","updated_at": "2167-08-14T05:39:34.78545966-04:00","deleted_at": "2038-04-28T16:05:30.980746651-04:00","miner": "uVwfePMZwYKPkYqvjBWNkOHhw","phase": "JOWXodRdKDwkeZOihORslIkWh","message": "iMBFYZGDPLDiEFxMFUDwqndJb","content": 73,"miner_version": "qeaWcudKiFPuLjhQlpjZBKorO","user_id": 74}' | http POST "http://localhost:3030/dferecords" X-Api-User:user123
func AddDfeRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	dferecords := &model.DfeRecord{}

	if err := readJSON(r, dferecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := dferecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	dferecords.Prepare()

	if err := dferecords.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "dfe_records", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	dferecords, _, err = dao.AddDfeRecords(ctx, dferecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, dferecords)
}

// UpdateDfeRecords Update a single record from dfe_records table in the estuary database
// @Summary Update an record in table dfe_records
// @Description Update a single record from dfe_records table in the estuary database
// @Tags DfeRecords
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  DfeRecords body model.DfeRecord true "Update DfeRecords record"
// @Success 200 {object} model.DfeRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dferecords/{argID} [put]
// echo '{"id": 27,"created_at": "2180-11-16T17:21:14.605962774-05:00","updated_at": "2167-08-14T05:39:34.78545966-04:00","deleted_at": "2038-04-28T16:05:30.980746651-04:00","miner": "uVwfePMZwYKPkYqvjBWNkOHhw","phase": "JOWXodRdKDwkeZOihORslIkWh","message": "iMBFYZGDPLDiEFxMFUDwqndJb","content": 73,"miner_version": "qeaWcudKiFPuLjhQlpjZBKorO","user_id": 74}' | http PUT "http://localhost:3030/dferecords/1"  X-Api-User:user123
func UpdateDfeRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	dferecords := &model.DfeRecord{}
	if err := readJSON(r, dferecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := dferecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	dferecords.Prepare()

	if err := dferecords.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "dfe_records", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	dferecords, _, err = dao.UpdateDfeRecords(ctx,
		argID,
		dferecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, dferecords)
}

// DeleteDfeRecords Delete a single record from dfe_records table in the estuary database
// @Summary Delete a record from dfe_records
// @Description Delete a single record from dfe_records table in the estuary database
// @Tags DfeRecords
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.DfeRecord
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dferecords/{argID} [delete]
// http DELETE "http://localhost:3030/dferecords/1" X-Api-User:user123
func DeleteDfeRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "dfe_records", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDfeRecords(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
