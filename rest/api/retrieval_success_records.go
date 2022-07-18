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

func configGinRetrievalSuccessRecordsRouter(router gin.IRoutes) {
	router.GET("/retrievalsuccessrecords", ConverHttpRouterToGin(GetAllRetrievalSuccessRecords))
	router.POST("/retrievalsuccessrecords", ConverHttpRouterToGin(AddRetrievalSuccessRecords))
	router.GET("/retrievalsuccessrecords/:argPropCid", ConverHttpRouterToGin(GetRetrievalSuccessRecords))
	router.PUT("/retrievalsuccessrecords/:argPropCid", ConverHttpRouterToGin(UpdateRetrievalSuccessRecords))
	router.DELETE("/retrievalsuccessrecords/:argPropCid", ConverHttpRouterToGin(DeleteRetrievalSuccessRecords))
}

// GetAllRetrievalSuccessRecords is a function to get a slice of record(s) from retrieval_success_records table in the estuary database
// @Summary Get list of RetrievalSuccessRecords
// @Tags RetrievalSuccessRecords
// @Description GetAllRetrievalSuccessRecords is a handler to get a slice of record(s) from retrieval_success_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RetrievalSuccessRecord}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /retrievalsuccessrecords [get]
// http "http://localhost:3030/retrievalsuccessrecords?page=0&pagesize=20" X-Api-User:user123
func GetAllRetrievalSuccessRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "retrieval_success_records", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRetrievalSuccessRecords(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRetrievalSuccessRecords is a function to get a single record from the retrieval_success_records table in the estuary database
// @Summary Get record from table RetrievalSuccessRecords by  argPropCid
// @Tags RetrievalSuccessRecords
// 
// @Description GetRetrievalSuccessRecords is a function to get a single record from the retrieval_success_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argPropCid path string true "prop_cid"
// @Success 200 {object} model.RetrievalSuccessRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /retrievalsuccessrecords/{argPropCid} [get]
// http "http://localhost:3030/retrievalsuccessrecords/hello world" X-Api-User:user123
func GetRetrievalSuccessRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPropCid, err := parseString(ps, "argPropCid")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_success_records", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRetrievalSuccessRecords(ctx, argPropCid)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRetrievalSuccessRecords add to add a single record to retrieval_success_records table in the estuary database
// @Summary Add an record to retrieval_success_records table
// @Description add to add a single record to retrieval_success_records table in the estuary database
// @Tags RetrievalSuccessRecords
// @Accept  json
// @Produce  json
// @Param RetrievalSuccessRecords body model.RetrievalSuccessRecord true "Add RetrievalSuccessRecords"
// @Success 200 {object} model.RetrievalSuccessRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /retrievalsuccessrecords [post]
// echo '{"prop_cid": "RWHHcumQyCbIYkXocrvWqpvOQ","miner": "FoAxSIpTSQLngfcwxRpGfSWoy","peer": "GnDVbyfNPhmOlLFwIxXYrGqUY","size": 24,"duration_ms": 89,"average_speed": 24,"total_payment": "RUWKLSoYqyNSEwkOBxRlZlTHk","num_payments": 28,"ask_price": "gHNGqkZAiEDSGPjhZvulTVyjf","id": 52,"created_at": "2149-09-11T05:35:47.126224067-04:00","cid": "rWkZJtNWKFxrjofwrUQQtMhKh"}' | http POST "http://localhost:3030/retrievalsuccessrecords" X-Api-User:user123
func AddRetrievalSuccessRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	retrievalsuccessrecords := &model.RetrievalSuccessRecord{}

	if err := readJSON(r, retrievalsuccessrecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := retrievalsuccessrecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	retrievalsuccessrecords.Prepare()

	if err := retrievalsuccessrecords.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_success_records", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	retrievalsuccessrecords, _, err = dao.AddRetrievalSuccessRecords(ctx, retrievalsuccessrecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, retrievalsuccessrecords)
}

// UpdateRetrievalSuccessRecords Update a single record from retrieval_success_records table in the estuary database
// @Summary Update an record in table retrieval_success_records
// @Description Update a single record from retrieval_success_records table in the estuary database
// @Tags RetrievalSuccessRecords
// @Accept  json
// @Produce  json
// @Param  argPropCid path string true "prop_cid"
// @Param  RetrievalSuccessRecords body model.RetrievalSuccessRecord true "Update RetrievalSuccessRecords record"
// @Success 200 {object} model.RetrievalSuccessRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /retrievalsuccessrecords/{argPropCid} [put]
// echo '{"prop_cid": "RWHHcumQyCbIYkXocrvWqpvOQ","miner": "FoAxSIpTSQLngfcwxRpGfSWoy","peer": "GnDVbyfNPhmOlLFwIxXYrGqUY","size": 24,"duration_ms": 89,"average_speed": 24,"total_payment": "RUWKLSoYqyNSEwkOBxRlZlTHk","num_payments": 28,"ask_price": "gHNGqkZAiEDSGPjhZvulTVyjf","id": 52,"created_at": "2149-09-11T05:35:47.126224067-04:00","cid": "rWkZJtNWKFxrjofwrUQQtMhKh"}' | http PUT "http://localhost:3030/retrievalsuccessrecords/hello world"  X-Api-User:user123
func UpdateRetrievalSuccessRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPropCid, err := parseString(ps, "argPropCid")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	retrievalsuccessrecords := &model.RetrievalSuccessRecord{}
	if err := readJSON(r, retrievalsuccessrecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := retrievalsuccessrecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	retrievalsuccessrecords.Prepare()

	if err := retrievalsuccessrecords.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_success_records", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	retrievalsuccessrecords, _, err = dao.UpdateRetrievalSuccessRecords(ctx,
		argPropCid,
		retrievalsuccessrecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, retrievalsuccessrecords)
}

// DeleteRetrievalSuccessRecords Delete a single record from retrieval_success_records table in the estuary database
// @Summary Delete a record from retrieval_success_records
// @Description Delete a single record from retrieval_success_records table in the estuary database
// @Tags RetrievalSuccessRecords
// @Accept  json
// @Produce  json
// @Param  argPropCid path string true "prop_cid"
// @Success 204 {object} model.RetrievalSuccessRecord
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /retrievalsuccessrecords/{argPropCid} [delete]
// http DELETE "http://localhost:3030/retrievalsuccessrecords/hello world" X-Api-User:user123
func DeleteRetrievalSuccessRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argPropCid, err := parseString(ps, "argPropCid")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_success_records", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRetrievalSuccessRecords(ctx, argPropCid)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
