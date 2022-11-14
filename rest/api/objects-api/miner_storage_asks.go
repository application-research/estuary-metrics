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

func ConfigMinerStorageAsksRouter(router gin.IRoutes) {
	router.GET("/minerstorageasks", api.ConvertHttpRouterToGin(GetAllMinerStorageAsks))
	router.GET("/minerstorageasks/:id", api.ConvertHttpRouterToGin(GetMinerStorageAsks))
}

// GetAllMinerStorageAsks is a function to get a slice of record(s) from miner_storage_asks table in the estuary database
// @Summary Get list of MinerStorageAsks
// @Tags MinerStorageAsks
// @Description GetAllMinerStorageAsks is a handler to get a slice of record(s) from miner_storage_asks table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.MinerStorageAsk}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minerstorageasks [get]
// http "http://localhost:3030/minerstorageasks?page=0&pagesize=20" X-Api-User:user123
func GetAllMinerStorageAsks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "miner_storage_asks", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllMinerStorageAsks(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetMinerStorageAsks is a function to get a single record from the miner_storage_asks table in the estuary database
// @Summary Get record from table MinerStorageAsks by  argID
// @Tags MinerStorageAsks
//
// @Description GetMinerStorageAsks is a function to get a single record from the miner_storage_asks table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.MinerStorageAsk
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /minerstorageasks/{argID} [get]
// http "http://localhost:3030/minerstorageasks/1" X-Api-User:user123
func GetMinerStorageAsks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "miner_storage_asks", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetMinerStorageAsks(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
