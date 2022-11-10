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

func ConfigPieceCommRecordsRouter(router gin.IRoutes) {
	router.GET("/piececommrecords", api.ConverHttpRouterToGin(GetAllPieceCommRecords))
	router.GET("/piececommrecords/:argData", api.ConverHttpRouterToGin(GetPieceCommRecords))
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
// @Success 200 {object} objects-api.PagedResults{data=[]model.PieceCommRecord}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /piececommrecords [get]
// http "http://localhost:3030/piececommrecords?page=0&pagesize=20" X-Api-User:user123
func GetAllPieceCommRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "piece_comm_records", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPieceCommRecords(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
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
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /piececommrecords/{argData} [get]
// http "http://localhost:3030/piececommrecords/hello world" X-Api-User:user123
func GetPieceCommRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argData, err := api.ParseString(ps, "argData")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "piece_comm_records", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPieceCommRecords(ctx, argData)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
