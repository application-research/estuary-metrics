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

func configGinContentDealsRouter(router gin.IRoutes) {
	router.GET("/contentdeals", ConverHttpRouterToGin(GetAllContentDeals))
	router.GET("/contentdeals/:argID", ConverHttpRouterToGin(GetContentDeals))
}

// GetAllContentDeals is a function to get a slice of record(s) from content_deals table in the estuary database
// @Summary Get list of ContentDeals
// @Tags ContentDeals
// @Description GetAllContentDeals is a handler to get a slice of record(s) from content_deals table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.ContentDeal}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdeals [get]
// http "http://localhost:3030/contentdeals?page=0&pagesize=20" X-Api-User:user123
func GetAllContentDeals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "content_deals", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentDeals(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetContentDeals is a function to get a single record from the content_deals table in the estuary database
// @Summary Get record from table ContentDeals by  argID
// @Tags ContentDeals
//
// @Description GetContentDeals is a function to get a single record from the content_deals table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.ContentDeal
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contentdeals/{argID} [get]
// http "http://localhost:3030/contentdeals/1" X-Api-User:user123
func GetContentDeals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deals", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentDeals(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddContentDeals add to add a single record to content_deals table in the estuary database
// @Summary Add an record to content_deals table
// @Description add to add a single record to content_deals table in the estuary database
// @Tags ContentDeals
// @Accept  json
// @Produce  json
// @Param ContentDeals body model.ContentDeal true "Add ContentDeals"
// @Success 200 {object} model.ContentDeal
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdeals [post]
// echo '{"id": 66,"created_at": "2117-02-22T05:39:44.052955639-05:00","updated_at": "2264-09-08T01:22:23.459000491-04:00","deleted_at": "2022-09-26T11:41:51.610003834-04:00","content": 85,"prop_cid": "yeinBhtFQwYVfKFRiVWLfRYfr","miner": "QUCGdjbHqCWOpdUUjIigiaRLh","deal_id": 81,"failed": false,"failed_at": "2305-08-10T11:18:08.502027902-04:00","dt_chan": "nTbFNNsgvGvlFBgPPNOgMmQrG","verified": false,"sealed_at": "2253-08-28T17:33:51.635839706-04:00","on_chain_at": "2059-08-20T21:08:18.804191574-04:00","transfer_started": "2246-10-22T02:49:37.887227108-04:00","transfer_finished": "2162-08-25T08:30:49.137350506-04:00","deal_uuid": "HupMissPrCUfqWYWrySmlVsZw","user_id": 96,"slashed": true}' | http POST "http://localhost:3030/contentdeals" X-Api-User:user123
func AddContentDeals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	contentdeals := &model.ContentDeal{}

	if err := readJSON(r, contentdeals); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdeals.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdeals.Prepare()

	if err := contentdeals.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deals", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	contentdeals, _, err = dao.AddContentDeals(ctx, contentdeals)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdeals)
}

// UpdateContentDeals Update a single record from content_deals table in the estuary database
// @Summary Update an record in table content_deals
// @Description Update a single record from content_deals table in the estuary database
// @Tags ContentDeals
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  ContentDeals body model.ContentDeal true "Update ContentDeals record"
// @Success 200 {object} model.ContentDeal
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contentdeals/{argID} [put]
// echo '{"id": 66,"created_at": "2117-02-22T05:39:44.052955639-05:00","updated_at": "2264-09-08T01:22:23.459000491-04:00","deleted_at": "2022-09-26T11:41:51.610003834-04:00","content": 85,"prop_cid": "yeinBhtFQwYVfKFRiVWLfRYfr","miner": "QUCGdjbHqCWOpdUUjIigiaRLh","deal_id": 81,"failed": false,"failed_at": "2305-08-10T11:18:08.502027902-04:00","dt_chan": "nTbFNNsgvGvlFBgPPNOgMmQrG","verified": false,"sealed_at": "2253-08-28T17:33:51.635839706-04:00","on_chain_at": "2059-08-20T21:08:18.804191574-04:00","transfer_started": "2246-10-22T02:49:37.887227108-04:00","transfer_finished": "2162-08-25T08:30:49.137350506-04:00","deal_uuid": "HupMissPrCUfqWYWrySmlVsZw","user_id": 96,"slashed": true}' | http PUT "http://localhost:3030/contentdeals/1"  X-Api-User:user123
func UpdateContentDeals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdeals := &model.ContentDeal{}
	if err := readJSON(r, contentdeals); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := contentdeals.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	contentdeals.Prepare()

	if err := contentdeals.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deals", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	contentdeals, _, err = dao.UpdateContentDeals(ctx,
		argID,
		contentdeals)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, contentdeals)
}

// DeleteContentDeals Delete a single record from content_deals table in the estuary database
// @Summary Delete a record from content_deals
// @Description Delete a single record from content_deals table in the estuary database
// @Tags ContentDeals
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.ContentDeal
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /contentdeals/{argID} [delete]
// http DELETE "http://localhost:3030/contentdeals/1" X-Api-User:user123
func DeleteContentDeals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "content_deals", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteContentDeals(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
