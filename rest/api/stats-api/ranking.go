package statsapi

import (
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ConfigRankingRoute(router gin.IRoutes) {
	router.GET("/rank/miners/:top", api.ConvertHttpRouterToGin(GetTopMiners))
	router.GET("/rank/users/:top", api.ConvertHttpRouterToGin(GetTopUsers))
	router.GET("/rank/users/collection/:top", api.ConvertHttpRouterToGin(GetTopCollectionUsers))
}

func GetTopMiners(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//dao.GetTopStorageMiners()

}

func GetTopUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//top := ps.ByName("top")
}

func GetTopCollectionUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//top := ps.ByName("top")
}
