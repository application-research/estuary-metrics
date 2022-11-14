package objectsapi

import (
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/application-research/estuary-metrics/rest/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func ConfigCollectionRefsRouter(router gin.IRoutes) {
	router.GET("/collectionrefs", api.ConvertHttpRouterToGin(GetAllCollectionRefs))
	router.GET("/collectionrefs/:id", api.ConvertHttpRouterToGin(GetCollectionRefs))
	router.GET("/collectionrefs/dynamicquery", api.ConvertHttpRouterToGin(GetCollectionsRefDynamicQuery))
}

// GetCollectionsRefDynamicQuery is a function to get a slice of record(s) from collection_refs table in the estuary database
// @Summary Get list of CollectionRefs
// @Tags CollectionRefs
// @Description GetCollectionsRefDynamicQuery is a handler to get a slice of record(s) from collection_refs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Param 	query query string false "query string"
// @Success 200 {object} api.PagedResults{data=[]model.CollectionRef}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /collectionrefs/dynamicquery [get]
func GetCollectionsRefDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	HandleDynamicQuery(w, r, ps, model.CollectionRef{})
}

// GetAllCollectionRefs is a function to get a slice of record(s) from collection_refs table in the estuary database
// @Summary Get list of CollectionRefs
// @Tags CollectionRefs
// @Description GetAllCollectionRefs is a handler to get a slice of record(s) from collection_refs table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.CollectionRef}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /collectionrefs [get]
// http "http://localhost:3030/collectionrefs?page=0&pagesize=20" X-Api-User:user123
func GetAllCollectionRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "collection_refs", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCollectionRefs(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetCollectionRefs is a function to get a single record from the collection_refs table in the estuary database
// @Summary Get record from table CollectionRefs by  argID
// @Tags CollectionRefs
//
// @Description GetCollectionRefs is a function to get a single record from the collection_refs table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.CollectionRef
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /collectionrefs/{argID} [get]
// http "http://localhost:3030/collectionrefs/1" X-Api-User:user123
func GetCollectionRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "collection_refs", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCollectionRefs(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
