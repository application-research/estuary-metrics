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

func configGinPieceCommRecordsRouter(router gin.IRoutes) {
	router.GET("/piececommrecords", ConverHttpRouterToGin(GetAllPieceCommRecords))
	router.GET("/piececommrecords/:argData", ConverHttpRouterToGin(GetPieceCommRecords))
}

// GetAllPieceCommRecords is a function to get a slice of record(s) from piece_comm_records table in the estuary database
// @Summary Get list of PieceCommRecords
// @Tags PieceCommRecords
// @Description GetAllPieceCommRecords is a handler to get a slice of record(s) from piece_comm_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.PieceCommRecord}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /piececommrecords [get]
// http "http://localhost:3030/piececommrecords?page=0&pagesize=20" X-Api-User:user123
func GetAllPieceCommRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "piece_comm_records", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPieceCommRecords(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetPieceCommRecords is a function to get a single record from the piece_comm_records table in the estuary database
// @Summary Get record from table PieceCommRecords by  argData
// @Tags PieceCommRecords
// @ID argData
// @Description GetPieceCommRecords is a function to get a single record from the piece_comm_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argData path string true "data"
// @Success 200 {object} model.PieceCommRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /piececommrecords/{argData} [get]
// http "http://localhost:3030/piececommrecords/hello world" X-Api-User:user123
func GetPieceCommRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argData, err := parseString(ps, "argData")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_comm_records", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPieceCommRecords(ctx, argData)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddPieceCommRecords add to add a single record to piece_comm_records table in the estuary database
// @Summary Add an record to piece_comm_records table
// @Description add to add a single record to piece_comm_records table in the estuary database
// @Tags PieceCommRecords
// @Accept  json
// @Produce  json
// @Param PieceCommRecords body model.PieceCommRecord true "Add PieceCommRecords"
// @Success 200 {object} model.PieceCommRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /piececommrecords [post]
// echo '{"data": "pfopAFpfNgxxkJpURKTtiBndD","piece": "EhKpQURAbBDQMEZNrwSvFKikV","size": 94,"car_size": 29}' | http POST "http://localhost:3030/piececommrecords" X-Api-User:user123
func AddPieceCommRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	piececommrecords := &model.PieceCommRecord{}

	if err := readJSON(r, piececommrecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := piececommrecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	piececommrecords.Prepare()

	if err := piececommrecords.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_comm_records", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	piececommrecords, _, err = dao.AddPieceCommRecords(ctx, piececommrecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, piececommrecords)
}

// UpdatePieceCommRecords Update a single record from piece_comm_records table in the estuary database
// @Summary Update an record in table piece_comm_records
// @Description Update a single record from piece_comm_records table in the estuary database
// @Tags PieceCommRecords
// @Accept  json
// @Produce  json
// @Param  argData path string true "data"
// @Param  PieceCommRecords body model.PieceCommRecord true "Update PieceCommRecords record"
// @Success 200 {object} model.PieceCommRecord
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /piececommrecords/{argData} [put]
// echo '{"data": "pfopAFpfNgxxkJpURKTtiBndD","piece": "EhKpQURAbBDQMEZNrwSvFKikV","size": 94,"car_size": 29}' | http PUT "http://localhost:3030/piececommrecords/hello world"  X-Api-User:user123
func UpdatePieceCommRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argData, err := parseString(ps, "argData")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	piececommrecords := &model.PieceCommRecord{}
	if err := readJSON(r, piececommrecords); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := piececommrecords.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	piececommrecords.Prepare()

	if err := piececommrecords.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_comm_records", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	piececommrecords, _, err = dao.UpdatePieceCommRecords(ctx,
		argData,
		piececommrecords)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, piececommrecords)
}

// DeletePieceCommRecords Delete a single record from piece_comm_records table in the estuary database
// @Summary Delete a record from piece_comm_records
// @Description Delete a single record from piece_comm_records table in the estuary database
// @Tags PieceCommRecords
// @Accept  json
// @Produce  json
// @Param  argData path string true "data"
// @Success 204 {object} model.PieceCommRecord
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /piececommrecords/{argData} [delete]
// http DELETE "http://localhost:3030/piececommrecords/hello world" X-Api-User:user123
func DeletePieceCommRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argData, err := parseString(ps, "argData")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "piece_comm_records", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeletePieceCommRecords(ctx, argData)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
