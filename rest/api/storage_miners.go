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

func configGinStorageMinersRouter(router gin.IRoutes) {
	router.GET("/storageminers", ConverHttpRouterToGin(GetAllStorageMiners))
	router.POST("/storageminers", ConverHttpRouterToGin(AddStorageMiners))
	router.GET("/storageminers/:argID", ConverHttpRouterToGin(GetStorageMiners))
	router.PUT("/storageminers/:argID", ConverHttpRouterToGin(UpdateStorageMiners))
	router.DELETE("/storageminers/:argID", ConverHttpRouterToGin(DeleteStorageMiners))
}

// GetAllStorageMiners is a function to get a slice of record(s) from storage_miners table in the estuary database
// @Summary Get list of StorageMiners
// @Tags StorageMiners
// @Description GetAllStorageMiners is a handler to get a slice of record(s) from storage_miners table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.StorageMiner}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /storageminers [get]
// http "http://localhost:3030/storageminers?page=0&pagesize=20" X-Api-User:user123
func GetAllStorageMiners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "storage_miners", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllStorageMiners(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetStorageMiners is a function to get a single record from the storage_miners table in the estuary database
// @Summary Get record from table StorageMiners by  argID
// @Tags StorageMiners
//
// @Description GetStorageMiners is a function to get a single record from the storage_miners table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.StorageMiner
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /storageminers/{argID} [get]
// http "http://localhost:3030/storageminers/1" X-Api-User:user123
func GetStorageMiners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "storage_miners", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetStorageMiners(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddStorageMiners add to add a single record to storage_miners table in the estuary database
// @Summary Add an record to storage_miners table
// @Description add to add a single record to storage_miners table in the estuary database
// @Tags StorageMiners
// @Accept  json
// @Produce  json
// @Param StorageMiners body model.StorageMiner true "Add StorageMiners"
// @Success 200 {object} model.StorageMiner
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /storageminers [post]
// echo '{"id": 1,"created_at": "2060-08-20T14:15:36.830552329-04:00","updated_at": "2254-05-03T11:44:20.046676275-04:00","deleted_at": "2180-12-03T08:40:23.229993399-05:00","address": "TtPQrggjsPTcmXHmPelFlXCpR","suspended": true,"suspended_reason": "qIGJyZBHIASdMgMmYQUeheiUn","name": "mlDVXBFSLFuouIoYoXFRqaqZX","version": "dFqKlaQCRbLlPTCmeRcBcdHtB","location": "wJPTEeVdECRkbWeQKMmxDLIjB","owner": 48}' | http POST "http://localhost:3030/storageminers" X-Api-User:user123
func AddStorageMiners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	storageminers := &model.StorageMiner{}

	if err := readJSON(r, storageminers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := storageminers.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	storageminers.Prepare()

	if err := storageminers.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "storage_miners", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	storageminers, _, err = dao.AddStorageMiners(ctx, storageminers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, storageminers)
}

// UpdateStorageMiners Update a single record from storage_miners table in the estuary database
// @Summary Update an record in table storage_miners
// @Description Update a single record from storage_miners table in the estuary database
// @Tags StorageMiners
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  StorageMiners body model.StorageMiner true "Update StorageMiners record"
// @Success 200 {object} model.StorageMiner
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /storageminers/{argID} [put]
// echo '{"id": 1,"created_at": "2060-08-20T14:15:36.830552329-04:00","updated_at": "2254-05-03T11:44:20.046676275-04:00","deleted_at": "2180-12-03T08:40:23.229993399-05:00","address": "TtPQrggjsPTcmXHmPelFlXCpR","suspended": true,"suspended_reason": "qIGJyZBHIASdMgMmYQUeheiUn","name": "mlDVXBFSLFuouIoYoXFRqaqZX","version": "dFqKlaQCRbLlPTCmeRcBcdHtB","location": "wJPTEeVdECRkbWeQKMmxDLIjB","owner": 48}' | http PUT "http://localhost:3030/storageminers/1"  X-Api-User:user123
func UpdateStorageMiners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	storageminers := &model.StorageMiner{}
	if err := readJSON(r, storageminers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := storageminers.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	storageminers.Prepare()

	if err := storageminers.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "storage_miners", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	storageminers, _, err = dao.UpdateStorageMiners(ctx,
		argID,
		storageminers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, storageminers)
}

// DeleteStorageMiners Delete a single record from storage_miners table in the estuary database
// @Summary Delete a record from storage_miners
// @Description Delete a single record from storage_miners table in the estuary database
// @Tags StorageMiners
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.StorageMiner
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /storageminers/{argID} [delete]
// http DELETE "http://localhost:3030/storageminers/1" X-Api-User:user123
func DeleteStorageMiners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "storage_miners", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteStorageMiners(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
