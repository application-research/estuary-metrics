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

// GetTopMiners is a function to get a slice of record(s) from miners table in the estuary database
// @Summary Get list of Miners
// @Description Get list of Miners
// @Tags Ranks
// @Accept  json
// @Produce  json
// @Param top path int true "top"
// @Success 200 {object} []dao.TopMiner
// @Failure 400 {object} api.HTTPError
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

//	GetTopUsers is a function to get a slice of record(s) from users table in the estuary database
//  @Summary Get list of Users
//  @Description Get list of Users
//  @Tags Ranks
//  @Accept  json
//  @Produce  json
//  @Param top path int true "top"
//  @Success 200 {object} []dao.TopUser
//  @Failure 400 {object} api.HTTPError
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

// GetTopCollectionUsers is a function to get a slice of record(s) from users table in the estuary database
// @Summary Get list of Users
// @Description Get list of Users
// @Tags Ranks
// @Accept  json
// @Produce  json
// @Param top path int true "top"
// @Success 200 {object} []dao.TopCollectionUser
// @Failure 400 {object} api.HTTPError
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
