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

func configGinMinerStorageAsksRouter(router gin.IRoutes) {
	router.GET("/minerstorageasks", ConverHttpRouterToGin(GetAllMinerStorageAsks))
	router.POST("/minerstorageasks", ConverHttpRouterToGin(AddMinerStorageAsks))
	router.GET("/minerstorageasks/:argID", ConverHttpRouterToGin(GetMinerStorageAsks))
	router.PUT("/minerstorageasks/:argID", ConverHttpRouterToGin(UpdateMinerStorageAsks))
	router.DELETE("/minerstorageasks/:argID", ConverHttpRouterToGin(DeleteMinerStorageAsks))
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

	if err := ValidateRequest(ctx, r, "miner_storage_asks", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllMinerStorageAsks(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
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
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "miner_storage_asks", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetMinerStorageAsks(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddMinerStorageAsks add to add a single record to miner_storage_asks table in the estuary database
// @Summary Add an record to miner_storage_asks table
// @Description add to add a single record to miner_storage_asks table in the estuary database
// @Tags MinerStorageAsks
// @Accept  json
// @Produce  json
// @Param MinerStorageAsks body model.MinerStorageAsk true "Add MinerStorageAsks"
// @Success 200 {object} model.MinerStorageAsk
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minerstorageasks [post]
// echo '{"id": 67,"created_at": "2224-08-04T06:32:39.509267727-04:00","updated_at": "2280-01-28T15:26:00.778411729-05:00","deleted_at": "2096-03-04T19:20:59.429825038-05:00","miner": "vgObBHuHBxYQtPgultYSLuCya","price": "WENCnIFayEKnMKjxCiEqTPwrB","verified_price": "sSAYRCSwgtlpEVUuftPqleUEU","min_piece_size": 77,"max_piece_size": 57}' | http POST "http://localhost:3030/minerstorageasks" X-Api-User:user123
func AddMinerStorageAsks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	minerstorageasks := &model.MinerStorageAsk{}

	if err := readJSON(r, minerstorageasks); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := minerstorageasks.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	minerstorageasks.Prepare()

	if err := minerstorageasks.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "miner_storage_asks", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	minerstorageasks, _, err = dao.AddMinerStorageAsks(ctx, minerstorageasks)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minerstorageasks)
}

// UpdateMinerStorageAsks Update a single record from miner_storage_asks table in the estuary database
// @Summary Update an record in table miner_storage_asks
// @Description Update a single record from miner_storage_asks table in the estuary database
// @Tags MinerStorageAsks
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  MinerStorageAsks body model.MinerStorageAsk true "Update MinerStorageAsks record"
// @Success 200 {object} model.MinerStorageAsk
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /minerstorageasks/{argID} [put]
// echo '{"id": 67,"created_at": "2224-08-04T06:32:39.509267727-04:00","updated_at": "2280-01-28T15:26:00.778411729-05:00","deleted_at": "2096-03-04T19:20:59.429825038-05:00","miner": "vgObBHuHBxYQtPgultYSLuCya","price": "WENCnIFayEKnMKjxCiEqTPwrB","verified_price": "sSAYRCSwgtlpEVUuftPqleUEU","min_piece_size": 77,"max_piece_size": 57}' | http PUT "http://localhost:3030/minerstorageasks/1"  X-Api-User:user123
func UpdateMinerStorageAsks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minerstorageasks := &model.MinerStorageAsk{}
	if err := readJSON(r, minerstorageasks); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := minerstorageasks.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	minerstorageasks.Prepare()

	if err := minerstorageasks.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "miner_storage_asks", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	minerstorageasks, _, err = dao.UpdateMinerStorageAsks(ctx,
		argID,
		minerstorageasks)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, minerstorageasks)
}

// DeleteMinerStorageAsks Delete a single record from miner_storage_asks table in the estuary database
// @Summary Delete a record from miner_storage_asks
// @Description Delete a single record from miner_storage_asks table in the estuary database
// @Tags MinerStorageAsks
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.MinerStorageAsk
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /minerstorageasks/{argID} [delete]
// http DELETE "http://localhost:3030/minerstorageasks/1" X-Api-User:user123
func DeleteMinerStorageAsks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "miner_storage_asks", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteMinerStorageAsks(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
