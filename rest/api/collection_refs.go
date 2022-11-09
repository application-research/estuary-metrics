package api

import (
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guregu/null"
	"github.com/julienschmidt/httprouter"
)

var (
	_ = null.Bool{}
)

func configGinCollectionRefsRouter(router gin.IRoutes) {
	router.GET("/collectionrefs", ConverHttpRouterToGin(GetAllCollectionRefs))
	router.GET("/collectionrefs/:argID", ConverHttpRouterToGin(GetCollectionRefs))
	router.DELETE("/collectionrefs/:argID", ConverHttpRouterToGin(DeleteCollectionRefs))
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

	if err := ValidateRequest(ctx, r, "collection_refs", model.RetrieveMany); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllCollectionRefs(ctx, int(page), int(pagesize), order)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	result := &PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	writeJSON(ctx, w, result)
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
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "collection_refs", model.RetrieveOne); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetCollectionRefs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, record)
}

// AddCollectionRefs add to add a single record to collection_refs table in the estuary database
// @Summary Add an record to collection_refs table
// @Description add to add a single record to collection_refs table in the estuary database
// @Tags CollectionRefs
// @Accept  json
// @Produce  json
// @Param CollectionRefs body model.CollectionRef true "Add CollectionRefs"
// @Success 200 {object} model.CollectionRef
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /collectionrefs [post]
// echo '{"id": 34,"created_at": "2128-07-19T10:28:43.238175479-04:00","collection": 99,"content": 22,"path": "uArqHOmoEOoEQyjaJggFWCAVy"}' | http POST "http://localhost:3030/collectionrefs" X-Api-User:user123
func AddCollectionRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)
	collectionrefs := &model.CollectionRef{}

	if err := readJSON(r, collectionrefs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := collectionrefs.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	collectionrefs.Prepare()

	if err := collectionrefs.Validate(model.Create); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "collection_refs", model.Create); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	var err error
	collectionrefs, _, err = dao.AddCollectionRefs(ctx, collectionrefs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, collectionrefs)
}

// UpdateCollectionRefs Update a single record from collection_refs table in the estuary database
// @Summary Update an record in table collection_refs
// @Description Update a single record from collection_refs table in the estuary database
// @Tags CollectionRefs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Param  CollectionRefs body model.CollectionRef true "Update CollectionRefs record"
// @Success 200 {object} model.CollectionRef
// @Failure 400 {object} api.HTTPError
// @Failure 404 {object} api.HTTPError
// @Router /collectionrefs/{argID} [put]
// echo '{"id": 34,"created_at": "2128-07-19T10:28:43.238175479-04:00","collection": 99,"content": 22,"path": "uArqHOmoEOoEQyjaJggFWCAVy"}' | http PUT "http://localhost:3030/collectionrefs/1"  X-Api-User:user123
func UpdateCollectionRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	collectionrefs := &model.CollectionRef{}
	if err := readJSON(r, collectionrefs); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := collectionrefs.BeforeSave(dao.DB); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
	}

	collectionrefs.Prepare()

	if err := collectionrefs.Validate(model.Update); err != nil {
		returnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := ValidateRequest(ctx, r, "collection_refs", model.Update); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	collectionrefs, _, err = dao.UpdateCollectionRefs(ctx,
		argID,
		collectionrefs)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeJSON(ctx, w, collectionrefs)
}

// DeleteCollectionRefs Delete a single record from collection_refs table in the estuary database
// @Summary Delete a record from collection_refs
// @Description Delete a single record from collection_refs table in the estuary database
// @Tags CollectionRefs
// @Accept  json
// @Produce  json
// @Param  argID path int64 true "id"
// @Success 204 {object} model.CollectionRef
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /collectionrefs/{argID} [delete]
// http DELETE "http://localhost:3030/collectionrefs/1" X-Api-User:user123
func DeleteCollectionRefs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := initializeContext(r)

	argID, err := parseInt64(ps, "argID")
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	if err := ValidateRequest(ctx, r, "collection_refs", model.Delete); err != nil {
		returnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteCollectionRefs(ctx, argID)
	if err != nil {
		returnError(ctx, w, r, err)
		return
	}

	writeRowsAffected(w, rowsAffected)
}
