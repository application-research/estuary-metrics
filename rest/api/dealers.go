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

func configGinDealersRouter(router gin.IRoutes) {
	router.GET("/dealers", ConverHttpRouterToGin(GetAllDealers))
	router.GET("/dealers/:argID", ConverHttpRouterToGin(GetDealers))
}

// GetAllDealers is a function to get a slice of record(s) from dealers table in the estuary database
// @Summary Get list of Dealers
// @Tags Dealers
// @Description GetAllDealers is a handler to get a slice of record(s) from dealers table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Dealer}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dealers [get]
// http "http://localhost:3030/dealers?page=0&pagesize=20" X-Api-User:user123
func GetAllDealers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "dealers", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDealers(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetDealers is a function to get a single record from the dealers table in the estuary database
// @Summary Get record from table Dealers by  argID
// @Tags Dealers
//
// @Description GetDealers is a function to get a single record from the dealers table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Dealer
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /dealers/{argID} [get]
// http "http://localhost:3030/dealers/1" X-Api-User:user123
func GetDealers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "dealers", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDealers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddDealers add to add a single record to dealers table in the estuary database
// @Summary Add an record to dealers table
// @Description add to add a single record to dealers table in the estuary database
// @Tags Dealers
// @Accept  json
// @Produce  json
// @Param Dealers body model.Dealer true "Add Dealers"
// @Success 200 {object} model.Dealer
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dealers [post]
// echo '{"id": 10,"created_at": "2136-06-29T08:44:35.542873193-04:00","updated_at": "2096-03-30T13:13:52.156618886-04:00","deleted_at": "2057-07-17T02:58:18.163426787-04:00","handle": "MqghfBSwtcQvPUXAyGQBBqXmu","token": "kwWvgcoFwciGIxksEnavEjwxr","host": "jdihUdSqOcFeBSoYbcEOUxNJC","peer_id": "RJgSNiuomVRKprfAyuJdOoYSy","open": false,"last_connection": "2312-11-02T04:19:49.875547768-04:00"}' | http POST "http://localhost:3030/dealers" X-Api-User:user123
func AddDealers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	dealers := &model.Dealer{}

	if err := readJSON(r, dealers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := dealers.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	dealers.Prepare()

	if err := dealers.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "dealers", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	dealers, _, err = dao.AddDealers(ctx, dealers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, dealers)
}

// UpdateDealers Update a single record from dealers table in the estuary database
// @Summary Update an record in table dealers
// @Description Update a single record from dealers table in the estuary database
// @Tags Dealers
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Dealers body model.Dealer true "Update Dealers record"
// @Success 200 {object} model.Dealer
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /dealers/{argID} [put]
// echo '{"id": 10,"created_at": "2136-06-29T08:44:35.542873193-04:00","updated_at": "2096-03-30T13:13:52.156618886-04:00","deleted_at": "2057-07-17T02:58:18.163426787-04:00","handle": "MqghfBSwtcQvPUXAyGQBBqXmu","token": "kwWvgcoFwciGIxksEnavEjwxr","host": "jdihUdSqOcFeBSoYbcEOUxNJC","peer_id": "RJgSNiuomVRKprfAyuJdOoYSy","open": false,"last_connection": "2312-11-02T04:19:49.875547768-04:00"}' | http PUT "http://localhost:3030/dealers/1"  X-Api-User:user123
func UpdateDealers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	dealers := &model.Dealer{}
	if err := readJSON(r, dealers); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := dealers.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	dealers.Prepare()

	if err := dealers.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "dealers", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	dealers, _, err = dao.UpdateDealers(ctx,
		argID,
		dealers)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, dealers)
}

// DeleteDealers Delete a single record from dealers table in the estuary database
// @Summary Delete a record from dealers
// @Description Delete a single record from dealers table in the estuary database
// @Tags Dealers
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.Dealer
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /dealers/{argID} [delete]
// http DELETE "http://localhost:3030/dealers/1" X-Api-User:user123
func DeleteDealers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "dealers", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteDealers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
