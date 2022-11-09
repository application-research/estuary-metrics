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

func configGinInviteCodesRouter(router gin.IRoutes) {
	router.GET("/invitecodes", ConverHttpRouterToGin(GetAllInviteCodes))
	router.GET("/invitecodes/:argID", ConverHttpRouterToGin(GetInviteCodes))
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
// @Success 200 {object} api.PagedResults{data=[]model.InviteCode}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /invitecodes [get]
// http "http://localhost:3030/invitecodes?page=0&pagesize=20" X-Api-User:user123
func GetAllInviteCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "invite_codes", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllInviteCodes(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
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
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /invitecodes/{argID} [get]
// http "http://localhost:3030/invitecodes/1" X-Api-User:user123
func GetInviteCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "invite_codes", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetInviteCodes(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddInviteCodes add to add a single record to invite_codes table in the estuary database
// @Summary Add an record to invite_codes table
// @Description add to add a single record to invite_codes table in the estuary database
// @Tags InviteCodes
// @Accept  json
// @Produce  json
// @Param InviteCodes body model.InviteCode true "Add InviteCodes"
// @Success 200 {object} model.InviteCode
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /invitecodes [post]
// echo '{"id": 45,"created_at": "2204-11-26T16:15:35.071092726-05:00","updated_at": "2190-02-11T21:23:44.90187849-05:00","deleted_at": "2213-06-16T13:14:13.911864722-04:00","code": "WsfvGGOuJyNRxUMhhSeqqZEVb","created_by": 80,"claimed_by": 93}' | http POST "http://localhost:3030/invitecodes" X-Api-User:user123
func AddInviteCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	invitecodes := &model.InviteCode{}

	if err := readJSON(r, invitecodes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := invitecodes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	invitecodes.Prepare()

	if err := invitecodes.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "invite_codes", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	invitecodes, _, err = dao.AddInviteCodes(ctx, invitecodes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, invitecodes)
}

// UpdateInviteCodes Update a single record from invite_codes table in the estuary database
// @Summary Update an record in table invite_codes
// @Description Update a single record from invite_codes table in the estuary database
// @Tags InviteCodes
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  InviteCodes body model.InviteCode true "Update InviteCodes record"
// @Success 200 {object} model.InviteCode
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /invitecodes/{argID} [put]
// echo '{"id": 45,"created_at": "2204-11-26T16:15:35.071092726-05:00","updated_at": "2190-02-11T21:23:44.90187849-05:00","deleted_at": "2213-06-16T13:14:13.911864722-04:00","code": "WsfvGGOuJyNRxUMhhSeqqZEVb","created_by": 80,"claimed_by": 93}' | http PUT "http://localhost:3030/invitecodes/1"  X-Api-User:user123
func UpdateInviteCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	invitecodes := &model.InviteCode{}
	if err := readJSON(r, invitecodes); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := invitecodes.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	invitecodes.Prepare()

	if err := invitecodes.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "invite_codes", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	invitecodes, _, err = dao.UpdateInviteCodes(ctx,
		argID,
		invitecodes)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, invitecodes)
}

// DeleteInviteCodes Delete a single record from invite_codes table in the estuary database
// @Summary Delete a record from invite_codes
// @Description Delete a single record from invite_codes table in the estuary database
// @Tags InviteCodes
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.InviteCode
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /invitecodes/{argID} [delete]
// http DELETE "http://localhost:3030/invitecodes/1" X-Api-User:user123
func DeleteInviteCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "invite_codes", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteInviteCodes(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
