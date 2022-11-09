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

func configGinRetrievalFailureRecordsRouter(router gin.IRoutes) {
	router.GET("/retrievalfailurerecords", ConverHttpRouterToGin(GetAllRetrievalFailureRecords))
	router.GET("/retrievalfailurerecords/:argID", ConverHttpRouterToGin(GetRetrievalFailureRecords))
}

// GetAllRetrievalFailureRecords is a function to get a slice of record(s) from retrieval_failure_records table in the estuary database
// @Summary Get list of RetrievalFailureRecords
// @Tags RetrievalFailureRecords
// @Description GetAllRetrievalFailureRecords is a handler to get a slice of record(s) from retrieval_failure_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.RetrievalFailureRecord}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /retrievalfailurerecords [get]
// http "http://localhost:3030/retrievalfailurerecords?page=0&pagesize=20" X-Api-User:user123
func GetAllRetrievalFailureRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "retrieval_failure_records", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRetrievalFailureRecords(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetRetrievalFailureRecords is a function to get a single record from the retrieval_failure_records table in the estuary database
// @Summary Get record from table RetrievalFailureRecords by  argID
// @Tags RetrievalFailureRecords
//
// @Description GetRetrievalFailureRecords is a function to get a single record from the retrieval_failure_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.RetrievalFailureRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /retrievalfailurerecords/{argID} [get]
// http "http://localhost:3030/retrievalfailurerecords/1" X-Api-User:user123
func GetRetrievalFailureRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_failure_records", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRetrievalFailureRecords(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddRetrievalFailureRecords add to add a single record to retrieval_failure_records table in the estuary database
// @Summary Add an record to retrieval_failure_records table
// @Description add to add a single record to retrieval_failure_records table in the estuary database
// @Tags RetrievalFailureRecords
// @Accept  json
// @Produce  json
// @Param RetrievalFailureRecords body model.RetrievalFailureRecord true "Add RetrievalFailureRecords"
// @Success 200 {object} model.RetrievalFailureRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /retrievalfailurerecords [post]
// echo '{"id": 22,"created_at": "2044-08-28T03:16:41.707263518-04:00","updated_at": "2164-12-20T08:43:10.867080355-05:00","deleted_at": "2122-10-26T10:18:20.284567107-04:00","miner": "IPAmhtbsaJSkxbwcxrxucgSZv","phase": "aYZXFdyGPLqIXfOOgMTTayIKZ","message": "CrtYmuIpNqWPGcjadDdjmvuTe","content": 44,"cid": "dgxDtYLafdZkcDyBipdsbFpND"}' | http POST "http://localhost:3030/retrievalfailurerecords" X-Api-User:user123
func AddRetrievalFailureRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	retrievalfailurerecords := &model.RetrievalFailureRecord{}

	if err := readJSON(r, retrievalfailurerecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := retrievalfailurerecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	retrievalfailurerecords.Prepare()

	if err := retrievalfailurerecords.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_failure_records", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	retrievalfailurerecords, _, err = dao.AddRetrievalFailureRecords(ctx, retrievalfailurerecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, retrievalfailurerecords)
}

// UpdateRetrievalFailureRecords Update a single record from retrieval_failure_records table in the estuary database
// @Summary Update an record in table retrieval_failure_records
// @Description Update a single record from retrieval_failure_records table in the estuary database
// @Tags RetrievalFailureRecords
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  RetrievalFailureRecords body model.RetrievalFailureRecord true "Update RetrievalFailureRecords record"
// @Success 200 {object} model.RetrievalFailureRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /retrievalfailurerecords/{argID} [put]
// echo '{"id": 22,"created_at": "2044-08-28T03:16:41.707263518-04:00","updated_at": "2164-12-20T08:43:10.867080355-05:00","deleted_at": "2122-10-26T10:18:20.284567107-04:00","miner": "IPAmhtbsaJSkxbwcxrxucgSZv","phase": "aYZXFdyGPLqIXfOOgMTTayIKZ","message": "CrtYmuIpNqWPGcjadDdjmvuTe","content": 44,"cid": "dgxDtYLafdZkcDyBipdsbFpND"}' | http PUT "http://localhost:3030/retrievalfailurerecords/1"  X-Api-User:user123
func UpdateRetrievalFailureRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	retrievalfailurerecords := &model.RetrievalFailureRecord{}
	if err := readJSON(r, retrievalfailurerecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := retrievalfailurerecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	retrievalfailurerecords.Prepare()

	if err := retrievalfailurerecords.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_failure_records", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	retrievalfailurerecords, _, err = dao.UpdateRetrievalFailureRecords(ctx,
		argID,
		retrievalfailurerecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, retrievalfailurerecords)
}

// DeleteRetrievalFailureRecords Delete a single record from retrieval_failure_records table in the estuary database
// @Summary Delete a record from retrieval_failure_records
// @Description Delete a single record from retrieval_failure_records table in the estuary database
// @Tags RetrievalFailureRecords
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.RetrievalFailureRecord
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /retrievalfailurerecords/{argID} [delete]
// http DELETE "http://localhost:3030/retrievalfailurerecords/1" X-Api-User:user123
func DeleteRetrievalFailureRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "retrieval_failure_records", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteRetrievalFailureRecords(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
