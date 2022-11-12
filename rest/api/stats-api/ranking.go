package statsapi

import (
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func ConfigRankingRoute(router gin.IRoutes) {
	router.GET("/rank/miners/:top", api.ConvertHttpRouterToGin(GetTopMiners))
	router.GET("/rank/users/:top", api.ConvertHttpRouterToGin(GetTopUsers))
	router.GET("/rank/users/collection/:top", api.ConvertHttpRouterToGin(GetTopCollectionUsers))
}

func GetTopMiners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	top := ps.ByName("top")
	topInt, err := strconv.Atoi(top)
	miners, err := dao.GetTopStorageMiners(ctx, topInt)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, miners)
}

func GetTopUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//top := ps.ByName("top")
	ctx := api.InitializeContext(r)
	top := ps.ByName("top")
	topInt, err := strconv.Atoi(top)
	users, err := dao.GetTopUsers(ctx, topInt)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, users)
}

func GetTopCollectionUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	top := ps.ByName("top")
	topInt, err := strconv.Atoi(top)
	users, err := dao.GetTopCollectionUsers(ctx, topInt)
	if err != nil {
		api.ReturnError(ctx, w, r, err)
		return
	}
	api.WriteJSON(ctx, w, users)
}
