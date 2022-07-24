package api

import (
	"encoding/json"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configGinUsersRouter(router gin.IRoutes) {
	router.GET("/users", ConverHttpRouterToGin(GetAllUsers))
	router.POST("/users", ConverHttpRouterToGin(AddUsers))
	router.POST("/users/dynamic", ConverHttpRouterToGin(GetUsersDynamicQuery))
	router.GET("/users/count", ConverHttpRouterToGin(GetNumberOfUsers))
	router.GET("/users/within-range", ConverHttpRouterToGin(GetNumberOfUsersWithinRange))
	router.GET("/users/:argID", ConverHttpRouterToGin(GetUsers))
	router.PUT("/users/:argID", ConverHttpRouterToGin(UpdateUsers))
	router.DELETE("/users/:argID", ConverHttpRouterToGin(DeleteUsers))
}

func GetUsersDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
	query := r.FormValue("query")

	var queryMap map[string]interface{}
	json.Unmarshal([]byte(query), &queryMap)

	if err := ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, err := dao.GetUsersDynamicQuery(ctx, queryMap, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: len(records)}

	writeJSON(ctx, w, result)
}

// GetNumberOfUsersWithinRange is a function to get the number of record(s) from users table in the estuary database
// @Summary Get number of Users within range
// @Tags Users
// @Description GetNumberOfUsersWithinRange is a handler to get the number of record(s) from users table in the estuary database
// @Accept  json
// @Produce  json
// @Success 200 {object} api.PagedResults{data=int64}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users/count [get]
// http "http://localhost:3030/users/count" X-Api-User:user123
func GetNumberOfUsersWithinRange(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	timeStart, errStart := time.Parse("2006-01-02", r.FormValue("start"))
	timeEnd, errEnd := time.Parse("2006-01-02", r.FormValue("end"))
	if errStart != nil || errEnd != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	users, err := dao.GetNumberOfUsersWithinRange(ctx, timeStart, timeEnd)

	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	type Result struct {
		Count int           `json:"count"`
		Users []*model.User `json:"users"`
	}

	writeJSON(ctx, w, Result{Count: len(users), Users: users})
}

// GetNumberOfUsers is a function to get the number of record(s) from users table in the estuary database
// @Summary Get number of Users
// @Tags Users
// @Description GetNumberOfUsers is a handler to get the number of record(s) from users table in the estuary database
// @Accept  json
// @Produce  json
// @Success 200 {object} api.PagedResults{data=int64}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users/count [get]
// http "http://localhost:3030/users/within-range" X-Api-User:user123
func GetNumberOfUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	if err := ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	totalRows, err := dao.GetNumberOfUsers(ctx)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, totalRows)
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
// @Success 200 {object} api.PagedResults{data=[]model.User}
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users [get]
// http "http://localhost:3030/users?page=0&pagesize=20" X-Api-User:user123
func GetAllUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := ValidateRequest(ctx, r, "users", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}
	result, err := dao.Cacher.Get("get/all/users", time.Minute*10,
		func() (interface{}, error) {
			records, totalRows, err := dao.GetAllUsers(ctx, int(page), int(pagesize), order)
			if err != nil {
				returnError(ctx, w, r, err)
			}

			result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
			return result, err
		},
	)

	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, result)
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
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /users/{argID} [get]
// http "http://localhost:3030/users/1" X-Api-User:user123
func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetUsers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddUsers add to add a single record to users table in the estuary database
// @Summary Add an record to users table
// @Description add to add a single record to users table in the estuary database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param Users body model.User true "Add Users"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users [post]
// echo '{"id": 74,"created_at": "2163-04-23T21:21:45.363153215-04:00","updated_at": "2133-07-15T20:19:59.64560316-04:00","deleted_at": "2289-03-20T22:48:39.139984501-04:00","uuid": "JwLfnhkSOfQHVdQZxuHOyMWCZ","username": "KNmOTkGKBSLIaigXXePBkxxlF","pass_hash": "rEIxHdjlfpBlROSUiOvRZILYE","user_email": "JVgFMpeloHiMWQKcaZnCFmtMQ","perm": 30,"flags": 96,"address": "HAnWoysxcjEuSYWZKATNTeSAt","storage_disabled": false,"d_id": "GXIMcsnVquQrAsahtPCCvEpjc"}' | http POST "http://localhost:3030/users" X-Api-User:user123
func AddUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	users := &model.User{}

	if err := readJSON(r, users); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := users.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	users.Prepare()

	if err := users.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	users, _, err = dao.AddUsers(ctx, users)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, users)
}

// UpdateUsers Update a single record from users table in the estuary database
// @Summary Update an record in table users
// @Description Update a single record from users table in the estuary database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  Users body model.User true "Update Users record"
// @Success 200 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /users/{argID} [put]
// echo '{"id": 74,"created_at": "2163-04-23T21:21:45.363153215-04:00","updated_at": "2133-07-15T20:19:59.64560316-04:00","deleted_at": "2289-03-20T22:48:39.139984501-04:00","uuid": "JwLfnhkSOfQHVdQZxuHOyMWCZ","username": "KNmOTkGKBSLIaigXXePBkxxlF","pass_hash": "rEIxHdjlfpBlROSUiOvRZILYE","user_email": "JVgFMpeloHiMWQKcaZnCFmtMQ","perm": 30,"flags": 96,"address": "HAnWoysxcjEuSYWZKATNTeSAt","storage_disabled": false,"d_id": "GXIMcsnVquQrAsahtPCCvEpjc"}' | http PUT "http://localhost:3030/users/1"  X-Api-User:user123
func UpdateUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	users := &model.User{}
	if err := readJSON(r, users); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := users.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	users.Prepare()

	if err := users.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	users, _, err = dao.UpdateUsers(ctx,
		argID,
		users)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, users)
}

// DeleteUsers Delete a single record from users table in the estuary database
// @Summary Delete a record from users
// @Description Delete a single record from users table in the estuary database
// @Tags Users
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.User
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /users/{argID} [delete]
// http DELETE "http://localhost:3030/users/1" X-Api-User:user123
func DeleteUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "users", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteUsers(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
