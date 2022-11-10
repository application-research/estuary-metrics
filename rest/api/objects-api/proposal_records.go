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

func ConfigProposalRecordsRouter(router gin.IRoutes) {
	router.GET("/proposalrecords", api.ConvertHttpRouterToGin(GetAllProposalRecords))
	router.GET("/proposalrecords/:argPropCid", api.ConvertHttpRouterToGin(GetProposalRecords))
}

// GetAllProposalRecords is a function to get a slice of record(s) from proposal_records table in the estuary database
// @Summary Get list of ProposalRecords
// @Tags ProposalRecords
// @Description GetAllProposalRecords is a handler to get a slice of record(s) from proposal_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param   page     query    int     false        "page requested (defaults to 0)"
// @Param   pagesize query    int     false        "number of records in a page  (defaults to 20)"
// @Param   order    query    string  false        "db sort order column"
// @Success 200 {object} objects-api.PagedResults{data=[]model.ProposalRecord}
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /proposalrecords [get]
// http "http://localhost:3030/proposalrecords?page=0&pagesize=20" X-Api-User:user123
func GetAllProposalRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

	if err := api.ValidateRequest(ctx, r, "proposal_records", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.GetAllProposalRecords(ctx, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}
	api.WriteJSON(ctx, w, result)
}

// GetProposalRecords is a function to get a single record from the proposal_records table in the estuary database
// @Summary Get record from table ProposalRecords by  argPropCid
// @Tags ProposalRecords
//
// @Description GetProposalRecords is a function to get a single record from the proposal_records table in the estuary database
// @Accept  json
// @Produce  json
// @Param  argPropCid path string true "prop_cid"
// @Success 200 {object} model.ProposalRecord
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError "ErrNotFound, db record for id not found - returns NotFound HTTP 404 not found error"
// @Router /proposalrecords/{argPropCid} [get]
// http "http://localhost:3030/proposalrecords/hello world" X-Api-User:user123
func GetProposalRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argPropCid, err := api.ParseString(ps, "argPropCid")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "proposal_records", model.RetrieveOne); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	record, err := dao.GetProposalRecords(ctx, argPropCid)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, record)
}

// AddProposalRecords add to add a single record to proposal_records table in the estuary database
// @Summary Add an record to proposal_records table
// @Description add to add a single record to proposal_records table in the estuary database
// @Tags ProposalRecords
// @Accept  json
// @Produce  json
// @Param ProposalRecords body model.ProposalRecord true "Add ProposalRecords"
// @Success 200 {object} model.ProposalRecord
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /proposalrecords [post]
// echo '{"prop_cid": "uvImofLDyVAaNglqdbNUeVqLO","data": "MowQanJEuGZJZpVqAPUxLYLSq"}' | http POST "http://localhost:3030/proposalrecords" X-Api-User:user123
func AddProposalRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	proposalrecords := &model.ProposalRecord{}

	if err := api.ReadJSON(r, proposalrecords); err != nil {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := proposalrecords.BeforeSave(dao.DB); err != nil {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
	}

	proposalrecords.Prepare()

	if err := proposalrecords.Validate(model.Create); err != nil {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := api.ValidateRequest(ctx, r, "proposal_records", model.Create); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	var err error
	proposalrecords, _, err = dao.AddProposalRecords(ctx, proposalrecords)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, proposalrecords)
}

// UpdateProposalRecords Update a single record from proposal_records table in the estuary database
// @Summary Update an record in table proposal_records
// @Description Update a single record from proposal_records table in the estuary database
// @Tags ProposalRecords
// @Accept  json
// @Produce  json
// @Param  argPropCid path string true "prop_cid"
// @Param  ProposalRecords body model.ProposalRecord true "Update ProposalRecords record"
// @Success 200 {object} model.ProposalRecord
// @Failure 400 {object} objects-api.HTTPError
// @Failure 404 {object} objects-api.HTTPError
// @Router /proposalrecords/{argPropCid} [put]
// echo '{"prop_cid": "uvImofLDyVAaNglqdbNUeVqLO","data": "MowQanJEuGZJZpVqAPUxLYLSq"}' | http PUT "http://localhost:3030/proposalrecords/hello world"  X-Api-User:user123
func UpdateProposalRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argPropCid, err := api.ParseString(ps, "argPropCid")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	proposalrecords := &model.ProposalRecord{}
	if err := api.ReadJSON(r, proposalrecords); err != nil {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := proposalrecords.BeforeSave(dao.DB); err != nil {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
	}

	proposalrecords.Prepare()

	if err := proposalrecords.Validate(model.Update); err != nil {
		api.ReturnError(ctx, w, r, dao.ErrBadParams)
		return
	}

	if err := api.ValidateRequest(ctx, r, "proposal_records", model.Update); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	proposalrecords, _, err = dao.UpdateProposalRecords(ctx,
		argPropCid,
		proposalrecords)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteJSON(ctx, w, proposalrecords)
}

// DeleteProposalRecords Delete a single record from proposal_records table in the estuary database
// @Summary Delete a record from proposal_records
// @Description Delete a single record from proposal_records table in the estuary database
// @Tags ProposalRecords
// @Accept  json
// @Produce  json
// @Param  argPropCid path string true "prop_cid"
// @Success 204 {object} model.ProposalRecord
// @Failure 400 {object} objects-api.HTTPError
// @Failure 500 {object} objects-api.HTTPError
// @Router /proposalrecords/{argPropCid} [delete]
// http DELETE "http://localhost:3030/proposalrecords/hello world" X-Api-User:user123
func DeleteProposalRecords(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)

	argPropCid, err := api.ParseString(ps, "argPropCid")
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	if err := api.ValidateRequest(ctx, r, "proposal_records", model.Delete); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	rowsAffected, err := dao.DeleteProposalRecords(ctx, argPropCid)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	api.WriteRowsAffected(w, rowsAffected)
}
