package objectsapi

import (
	"encoding/json"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/application-research/estuary-metrics/rest/api"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func ConfigUsersRouter(router gin.IRoutes) {
	router.GET("/users", api.ConvertHttpRouterToGin(GetAllUsers))
	router.GET("/users/count", api.ConvertHttpRouterToGin(GetNumberOfUsers))
	router.GET("/users/within-range", api.ConvertHttpRouterToGin(GetNumberOfUsersWithinRange))
	router.GET("/users/:argID", api.ConvertHttpRouterToGin(GetUsers))
}

func GetUsersDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	query := r.FormValue("query")

	var queryMap map[string]interface{}
	json.Unmarshal([]byte(query), &queryMap)

	if err := api.ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, err := dao.GetUsersDynamicQuery(ctx, queryMap, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: len(records)}

	api.WriteJSON(ctx, w, result)
}

// GetNumberOfUsersWithinRange is a function to get the number of record(s) from users table in the estuary database
// @Summary Get number of Users within range
// @Tags Users
// @Description GetNumberOfUsersWithinRange is a handler to get the number of record(s) from users table in the estuary database
// @Accept  json
// @Produce  json
// @Success 200 {object} objects-api.PagedResults{data=int64}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /users/count [get]
// http "http://localhost:3030/users/count" X-Api-User:user123
func GetNumberOfUsersWithinRange(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	if err := api.ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	timeStart, errStart := time.Parse("2006-01-02", r.FormValue("start"))
	timeEnd, errEnd := time.Parse("2006-01-02", r.FormValue("end"))
	if errStart != nil || errEnd != nil {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	users, err := dao.GetNumberOfUsersWithinRange(ctx, timeStart, timeEnd)

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	type Result struct {
		Count int           `json:"count"`
		Users []*model.User `json:"users"`
	}

	api.WriteJSON(ctx, w, Result{Count: len(users), Users: users})
}

// GetNumberOfUsers is a function to get the number of record(s) from users table in the estuary database
// @Summary Get number of Users
// @Tags Users
// @Description GetNumberOfUsers is a handler to get the number of record(s) from users table in the estuary database
// @Accept  json
// @Produce  json
// @Success 200 {object} objects-api.PagedResults{data=int64}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /users/count [get]
// http "http://localhost:3030/users/within-range" X-Api-User:user123
func GetNumberOfUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	if err := api.ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	totalRows, err := dao.GetNumberOfUsers(ctx)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, totalRows)
}

// GetAllUsers is a function to get a slice of record(s) from users table in the estuary database
// @Summary Get list of Users
// @Tags Users
// @Description GetAllUsers is a handler to get a slice of record(s) from users table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} objects-api.PagedResults{data=[]model.User}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /users [get]
// http "http://localhost:3030/users?page=0&pagesize=20" X-Api-User:user123
func GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	result, err := dao.Cacher.Get("get/all/users", time.Minute*10,
		func() (interface{}, error) {
			records, totalRows, err := dao.GetAllUsers(ctx, int(page), int(pagesize), order)
			if err != nil {
				api.ReturnError(ctx, w, r, err)
			}

			result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
			return result, err
		},
	)

	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, result)
}

// GetUsers is a function to get a single record from the users table in the estuary database
// @Summary Get record from table Users by  argID
// @Tags Users
//
// @Description GetUsers is a function to get a single record from the users table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 200 {object} model.User
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /users/{argID} [get]
// http "http://localhost:3030/users/1" X-Api-User:user123
func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argID, err := api.ParseInt64(ps, "argID")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "users", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUsers(ctx, argID)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}
