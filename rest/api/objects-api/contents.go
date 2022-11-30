package objectsapi

import (
	"github.com/application-research/estuary-metrics/rest/api"
	"net/http"
	"time"

	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func ConfigContentsRouter(router gin.IRoutes) {
	router.GET("/contents", api.ConvertHttpRouterToGin(GetAllContents))
	router.GET("/contents/:id", api.ConvertHttpRouterToGin(GetContents))
	router.GET("/contents/dynamicquery", api.ConvertHttpRouterToGin(GetContentsDynamicQuery))
}

func ConfigUnProtectedContentsRouter(router gin.IRoutes) {
	router.GET("/contents/month-to-month/:months", api.ConvertHttpRouterToGin(AllContentOverThePastMonths))
	router.GET("/contents/set-months/:from/:to", api.ConvertHttpRouterToGin(AllContentOverASpecificDates))
}

func GetContentsDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HandleDynamicQuery(w, r, ps, model.Content{})
}

// GetAllContents is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description GetAllContents is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Content}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /contents [get]
// http "http://localhost:3030/contents?page=0&pagesize=20" X-Api-User:user123
func GetAllContents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "contents", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContents(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetContents is a function to get a single record from the contents table in the estuary database
// @Summary Get record from table Contents by  argID
// @Tags Contents
//
// @Description GetContents is a function to get a single record from the contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Content
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /contents/{argID} [get]
// http "http://localhost:3030/contents/1" X-Api-User:user123
func GetContents(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "contents", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.Cacher.Get("getContents", time.Minute*2, func() (interface{}, error) {
		return dao.GetContents(ctx, argID)
	})

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// AllContentOverThePastMonths is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description AllContentOverThePastMonths is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   months     query    int     false        "previous number of months"
func AllContentOverThePastMonths(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "months")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "contents", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.Cacher.Get("getContentMonthByMonth", time.Minute*2, func() (interface{}, error) {
		return dao.AllDataOverThePastMonth(ctx, model.Content{}, argID)
	})

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// AllContentOverASpecificDates	is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description AllContentOverASpecificDates is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   start     query    string     false        "start date"
// @Param   end     query    string     false        "end date"
func AllContentOverASpecificDates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)

	fromDate, err := api.ParseString(ps, "from")
	toDate, err := api.ParseString(ps, "to")

	record, err := dao.Cacher.Get("getContentOverASpecificDates", time.Minute*2, func() (interface{}, error) {
		return dao.AllDataOverASpecificDates(ctx, model.Content{}, fromDate, toDate)
	})

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
