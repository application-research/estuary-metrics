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

func ConfigDealersRouter(router gin.IRoutes) {
	router.GET("/dealers", api.ConvertHttpRouterToGin(GetAllDealers))
	router.GET("/dealers/:id", api.ConvertHttpRouterToGin(GetDealers))
	router.GET("/dealers/dynamicquery", api.ConvertHttpRouterToGin(GetDealersDynamicQuery))
}

func GetDealersDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HandleDynamicQuery(w, r, ps, model.Dealer{})
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

	if err := api.ValidateRequest(ctx, r, "dealers", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllDealers(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
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
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "dealers", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetDealers(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
