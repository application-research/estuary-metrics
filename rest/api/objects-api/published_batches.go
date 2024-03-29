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

func ConfigPublishedBatchesRouter(router gin.IRoutes) {
	router.GET("/publishedbatches", api.ConvertHttpRouterToGin(GetAllPublishedBatches))
	router.GET("/publishedbatches/:id", api.ConvertHttpRouterToGin(GetPublishedBatches))
}
func ConfigPublishedBatchesUnProtectedRouter(router gin.IRoutes) {
	router.GET("/publishedbatches/total", api.ConvertHttpRouterToGin(GetPublishedBatchesTotal))
}

// GetAllPublishedBatches is a function to get a slice of record(s) from published_batches table in the estuary database
// @Summary Get list of PublishedBatches
// @Description Get list of PublishedBatches
// @Tags PublishedBatches
// @Accept  json
// @Produce  json
// @Param page query int false "page"
// @Param pagesize query int false "pagesize"
// @Param order query string false "order"
// @Success 200 {object} api.PagedResults{data=[]PublishedBatches}
// @Failure 400 {object} api.Error
// @Router /publishedbatches [get]
func GetAllPublishedBatches(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "published_batches", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllPublishedBatches(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetPublishedBatches is a function to get a single record to published_batches table in the estuary database
// @Summary Get a PublishedBatches record
// @Description Get a PublishedBatches record
// @Tags PublishedBatches
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} PublishedBatches
// @Failure 400 {object} api.Error
// @Router /publishedbatches/{id} [get]
func GetPublishedBatches(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "published_batches", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetPublishedBatches(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// GetPublishedBatchesTotal is a function to get a single record to published_batches table in the estuary database
// @Summary Get a PublishedBatches record
// @Description Get a PublishedBatches record
// @Tags PublishedBatches
// @Accept  json
// @Produce  json
// @Success 200 {object} PublishedBatchesCount
// @Failure 400 {object} api.Error
// @Router /publishedbatches/total [get]
func GetPublishedBatchesTotal(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	if err := api.ValidateRequest(ctx, r, "published_batches", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetTotalPublishedBatches(ctx)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
