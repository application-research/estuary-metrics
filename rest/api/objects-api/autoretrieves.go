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

func ConfigAutoretrievesRouter(router gin.IRoutes) {
	router.GET("/autoretrieves", api.ConvertHttpRouterToGin(GetAllAutoretrieves))
	router.GET("/autoretrieves/:id", api.ConvertHttpRouterToGin(GetAutoretrieves))
	router.GET("/autoretrieves/dynamicquery", api.ConvertHttpRouterToGin(GetAutoretrievesDynamicQuery))
}

func ConfigUnProtectedAutoRetrievesRouter(router gin.IRoutes) {
	router.GET("/autoretrieves/month-to-month/:months", api.ConvertHttpRouterToGin(AllAutoRetrieveOverThePastMonths))
	router.GET("/autoretrieves/set-months/:from/:to", api.ConvertHttpRouterToGin(AllAutoRetrieveOverASpecificDates))
}

// GetAutoretrievesDynamicQuery is a function to get a slice of record(s) from autoretrieves table in the estuary database
// @Summary Get list of Autoretrieves
// @Tags Autoretrieves
// @Description GetAutoretrievesDynamicQuery is a function to get a slice of record(s) from autoretrieves table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Param   query    query    string  false        "dynamic query"
// @Success 200 {object} api.PagedResults{data=[]model.Autoretrieve}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /autoretrieves/dynamicquery [get]
func GetAutoretrievesDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HandleDynamicQuery(w, r, ps, model.AuthToken{})
}

// GetAllAutoretrieves is a function to get a slice of record(s) from autoretrieves table in the estuary database
// @Summary Get list of Autoretrieves
// @Tags Autoretrieves
// @Description GetAllAutoretrieves is a handler to get a slice of record(s) from autoretrieves table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.Autoretrieve}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /autoretrieves [get]
func GetAllAutoretrieves(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "autoretrieves", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAutoretrieves(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetAutoretrieves is a function to get a single record from the autoretrieves table in the estuary database
// @Summary Get record from table Autoretrieves by  argID
// @Tags Autoretrieves
//
// @Description GetAutoretrieves is a function to get a single record from the autoretrieves table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.Autoretrieve
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /autoretrieves/{argID} [get]
// http "http://localhost:3030/autoretrieves/1" X-Api-User:user123
func GetAutoretrieves(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "autoretrieves", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAutoretrieves(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// AllAutoRetrieveOverThePastMonths is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description AllContentOverThePastMonths is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   months     query    int     false        "previous number of months"
func AllAutoRetrieveOverThePastMonths(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "months")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.Cacher.Get("AllAutoRetrieveOverThePastMonths", time.Minute*2, func() (interface{}, error) {
		return dao.AllDataOverThePastMonth(ctx, model.Autoretrieve{}, argID)
	})

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// AllAutoRetrieveOverASpecificDates	is a function to get a slice of record(s) from contents table in the estuary database
// @Summary Get list of Contents
// @Tags Contents
// @Description AllContentOverASpecificDates is a handler to get a slice of record(s) from contents table in the estuary database
// @Accept  json
// @Produce  json
// @Param   start     query    string     false        "start date"
// @Param   end     query    string     false        "end date"
func AllAutoRetrieveOverASpecificDates(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctx := api.InitializeContext(r)

	fromDate, err := api.ParseString(ps, "from")
	toDate, err := api.ParseString(ps, "to")

	record, err := dao.Cacher.Get("AllAutoRetrieveOverASpecificDates", time.Minute*2, func() (interface{}, error) {
		return dao.AllDataOverASpecificDates(ctx, model.Autoretrieve{}, fromDate, toDate)
	})

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
