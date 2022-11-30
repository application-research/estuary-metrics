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

func ConfigContentDealsRouter(router gin.IRoutes) {
	router.GET("/contentdeals", api.ConvertHttpRouterToGin(GetAllContentDeals))
	router.GET("/contentdeals/:id", api.ConvertHttpRouterToGin(GetContentDeals))
	router.GET("/contentdeals/dynamicquery", api.ConvertHttpRouterToGin(GetContentDealsDynamicQuery))
}

func ConfigUnProtectedContentDealsRouter(router gin.IRoutes) {
	router.GET("/contentdeals/month-to-month/:months", api.ConvertHttpRouterToGin(AllContentDealsOverThePastMonths))
	router.GET("/contentdeals/set-months/:from/:to", api.ConvertHttpRouterToGin(AllContentDealsOverASpecificDates))
}

// GetContentDealsDynamicQuery is a function to get a slice of record(s) from content_deals table in the estuary database
// @Summary Get list of ContentDeals
// @Tags ContentDeals
// @Description GetContentDealsDynamicQuery is a handler to get a slice of record(s) from content_deals table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Param   query    query    string  false        "dynamic query"
// @Success 200 {object} api.PagedResults{data=[]model.ContentDeal}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
func GetContentDealsDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HandleDynamicQuery(w, r, ps, model.ContentDeal{})
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

	if err := api.ValidateRequest(ctx, r, "content_deals", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllContentDeals(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
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
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "content_deals", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetContentDeals(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// AllContentDealsOverThePastMonths is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description AllContentOverThePastMonths is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   months     query    int     false        "previous number of months"
func AllContentDealsOverThePastMonths(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "months")
	cache, err := api.ParseString(ps, "cache")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "contents", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	var record interface{}
	if cache == "y" {
		record, err = dao.Cacher.Get("getContentOverASpecificDates", time.Minute*2, func() (interface{}, error) {
			return dao.AllDataOverThePastMonth(ctx, model.ContentDeal{}, argID)
		})
	} else {
		record, err = dao.AllDataOverThePastMonth(ctx, model.ContentDeal{}, argID)
	}

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// AllContentDealsOverASpecificDates	is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description AllContentOverASpecificDates is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   start     query    string     false        "start date"
// @Param   end     query    string     false        "end date"
func AllContentDealsOverASpecificDates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)

	fromDate, err := api.ParseString(ps, "from")
	toDate, err := api.ParseString(ps, "to")
	cache, err := api.ParseString(ps, "cache")

	var record interface{}
	if cache == "y" {
		record, err = dao.Cacher.Get("getContentOverASpecificDates", time.Minute*2, func() (interface{}, error) {
			return dao.AllDataOverASpecificDates(ctx, model.ContentDeal{}, fromDate, toDate)
		})
	} else {
		record, err = dao.AllDataOverASpecificDates(ctx, model.ContentDeal{}, fromDate, toDate)
	}

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
