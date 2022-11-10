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

func ConfigInviteCodesRouter(router gin.IRoutes) {
	router.GET("/invitecodes", api.ConverHttpRouterToGin(GetAllInviteCodes))
	router.GET("/invitecodes/:argID", api.ConverHttpRouterToGin(GetInviteCodes))
}

// GetAllInviteCodes is a function to get a slice of record(s) from invite_codes table in the estuary database
// @Summary Get list of InviteCodes
// @Tags InviteCodes
// @Description GetAllInviteCodes is a handler to get a slice of record(s) from invite_codes table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} objects-api.PagedResults{data=[]model.InviteCode}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /invitecodes [get]
// http "http://localhost:3030/invitecodes?page=0&pagesize=20" X-Api-User:user123
func GetAllInviteCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "invite_codes", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllInviteCodes(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetInviteCodes is a function to get a single record from the invite_codes table in the estuary database
// @Summary Get record from table InviteCodes by  argID
// @Tags InviteCodes
//
// @Description GetInviteCodes is a function to get a single record from the invite_codes table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.InviteCode
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /invitecodes/{argID} [get]
// http "http://localhost:3030/invitecodes/1" X-Api-User:user123
func GetInviteCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "invite_codes", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetInviteCodes(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
