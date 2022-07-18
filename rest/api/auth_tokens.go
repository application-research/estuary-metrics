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

func configGinAuthTokensRouter(router gin.IRoutes) {
	router.GET("/authtokens", ConverHttpRouterToGin(GetAllAuthTokens))
	router.POST("/authtokens", ConverHttpRouterToGin(AddAuthTokens))
	router.GET("/authtokens/:argID", ConverHttpRouterToGin(GetAuthTokens))
	router.PUT("/authtokens/:argID", ConverHttpRouterToGin(UpdateAuthTokens))
	router.DELETE("/authtokens/:argID", ConverHttpRouterToGin(DeleteAuthTokens))
}

// GetAllAuthTokens is a function to get a slice of record(s) from auth_tokens table in the estuary database
// @Summary Get list of AuthTokens
// @Tags AuthTokens
// @Description GetAllAuthTokens is a handler to get a slice of record(s) from auth_tokens table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} api.PagedResults{data=[]model.AuthToken}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authtokens [get]
// http "http://localhost:3030/authtokens?page=0&pagesize=20" X-Api-User:user123
func GetAllAuthTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "auth_tokens", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllAuthTokens(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
}

// GetAuthTokens is a function to get a single record from the auth_tokens table in the estuary database
// @Summary Get record from table AuthTokens by  argID
// @Tags AuthTokens
//
// @Description GetAuthTokens is a function to get a single record from the auth_tokens table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.AuthToken
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /authtokens/{argID} [get]
// http "http://localhost:3030/authtokens/1" X-Api-User:user123
func GetAuthTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "auth_tokens", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetAuthTokens(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddAuthTokens add to add a single record to auth_tokens table in the estuary database
// @Summary Add an record to auth_tokens table
// @Description add to add a single record to auth_tokens table in the estuary database
// @Tags AuthTokens
// @Accept  json
// @Produce  json
// @Param AuthTokens body model.AuthToken true "Add AuthTokens"
// @Success 200 {object} model.AuthToken
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authtokens [post]
// echo '{"id": 97,"created_at": "2029-08-09T04:11:47.190457664-04:00","updated_at": "2092-05-04T13:53:01.735230585-04:00","deleted_at": "2183-05-17T20:59:59.686686765-04:00","token": "XekYoaHlOvwrLmWjBHFToIRTD","user": 17,"expiry": "2126-10-06T17:36:42.191060564-04:00","upload_only": false}' | http POST "http://localhost:3030/authtokens" X-Api-User:user123
func AddAuthTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	authtokens := &model.AuthToken{}

	if err := readJSON(r, authtokens); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authtokens.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authtokens.Prepare()

	if err := authtokens.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "auth_tokens", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	authtokens, _, err = dao.AddAuthTokens(ctx, authtokens)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authtokens)
}

// UpdateAuthTokens Update a single record from auth_tokens table in the estuary database
// @Summary Update an record in table auth_tokens
// @Description Update a single record from auth_tokens table in the estuary database
// @Tags AuthTokens
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  AuthTokens body model.AuthToken true "Update AuthTokens record"
// @Success 200 {object} model.AuthToken
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /authtokens/{argID} [put]
// echo '{"id": 97,"created_at": "2029-08-09T04:11:47.190457664-04:00","updated_at": "2092-05-04T13:53:01.735230585-04:00","deleted_at": "2183-05-17T20:59:59.686686765-04:00","token": "XekYoaHlOvwrLmWjBHFToIRTD","user": 17,"expiry": "2126-10-06T17:36:42.191060564-04:00","upload_only": false}' | http PUT "http://localhost:3030/authtokens/1"  X-Api-User:user123
func UpdateAuthTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authtokens := &model.AuthToken{}
	if err := readJSON(r, authtokens); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := authtokens.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	authtokens.Prepare()

	if err := authtokens.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "auth_tokens", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	authtokens, _, err = dao.UpdateAuthTokens(ctx,
		argID,
		authtokens)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, authtokens)
}

// DeleteAuthTokens Delete a single record from auth_tokens table in the estuary database
// @Summary Delete a record from auth_tokens
// @Description Delete a single record from auth_tokens table in the estuary database
// @Tags AuthTokens
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.AuthToken
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /authtokens/{argID} [delete]
// http DELETE "http://localhost:3030/authtokens/1" X-Api-User:user123
func DeleteAuthTokens(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "auth_tokens", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteAuthTokens(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
