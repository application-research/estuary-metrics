package objectsapi

import (
	"github.com/application-research/estuary-metrics/rest/api"
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

func ConfigRetrievalSuccessRecordsRouter(router gin.IRoutes) {
	router.GET("/retrievalsuccessrecords", api.ConvertHttpRouterToGin(GetAllRetrievalSuccessRecords))
	router.GET("/retrievalsuccessrecords/:argPropCid", api.ConvertHttpRouterToGin(GetRetrievalSuccessRecords))
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
// @Success 200 {object} objects-api.PagedResults{data=[]model.RetrievalSuccessRecord}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /retrievalsuccessrecords [get]
// http "http://localhost:3030/retrievalsuccessrecords?page=0&pagesize=20" X-Api-User:user123
func GetAllRetrievalSuccessRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	page, err := api.ReadInt(r, "page", 0)
	if err != nil || page < 0 {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	pagesize, err := api.ReadInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	order := r.FormValue("order")

	if err := api.ValidateRequest(ctx, r, "retrieval_success_records", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllRetrievalSuccessRecords(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
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
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /retrievalsuccessrecords/{argPropCid} [get]
// http "http://localhost:3030/retrievalsuccessrecords/hello world" X-Api-User:user123
func GetRetrievalSuccessRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argPropCid, err := api.ParseString(ps, "argPropCid")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "retrieval_success_records", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetRetrievalSuccessRecords(ctx, argPropCid)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
