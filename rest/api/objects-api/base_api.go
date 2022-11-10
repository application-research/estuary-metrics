package objectsapi

import (
	"encoding/json"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func HandleDynamicQuery(w http.ResponseWriter, r *http.Request, ps httprouter.Params, modelToQuery interface{}) {
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

	if err := api.ValidateRequest(ctx, r, "auth_tokens", model.RetrieveMany); err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	records, totalRows, err := dao.RunDynamicQuery(ctx, modelToQuery, queryMap, int(page), int(pagesize), order)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	result := &api.PagedResults{Page: page, PageSize: pagesize, Data: records, TotalRecords: int(totalRows)}

	api.WriteJSON(ctx, w, result)
}
