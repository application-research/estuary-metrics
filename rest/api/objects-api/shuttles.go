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

func ConfigShuttlesRouter(router gin.IRoutes) {
	router.GET("/shuttles", api.ConverHttpRouterToGin(GetAllShuttles))
	router.GET("/shuttles/count", api.ConverHttpRouterToGin(GetNumberOfRegisteredShuttles))
	router.GET("/shuttles/:argID", api.ConverHttpRouterToGin(GetShuttles))
}

func GetNumberOfRegisteredShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	if err := api.ValidateRequest(ctx, r, "shuttles", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	var totalRows int64
	dao.DB.Model(&model.Shuttle{}).Count(&totalRows)

	api.WriteJSON(ctx, w, totalRows)
}

// GetAllShuttles is a function to get a slice of record(s) from shuttles table in the estuary database
// @Summary Get list of Shuttles
// @Tags Shuttles
// @Description GetAllShuttles is a handler to get a slice of record(s) from shuttles table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} objects-api.PagedResults{data=[]model.Shuttle}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /shuttles [get]
// http "http://localhost:3030/shuttles?page=0&pagesize=20" X-Api-User:user123
func GetAllShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "shuttles", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllShuttles(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetShuttles is a function to get a single record from the shuttles table in the estuary database
// @Summary Get record from table Shuttles by  argID
// @Tags Shuttles
//
// @Description GetShuttles is a function to get a single record from the shuttles table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Shuttle
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /shuttles/{argID} [get]
// http "http://localhost:3030/shuttles/1" X-Api-User:user123
func GetShuttles(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "shuttles", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetShuttles(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
