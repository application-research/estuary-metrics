package statsapi

import (
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ConfigStatsRouter(router gin.IRoutes) {
	router.GET("/stats/retrieval", api.ConvertHttpRouterToGin(GetRetrievalStats))
	router.GET("/stats/storage", api.ConvertHttpRouterToGin(GetStorageStats))
	router.GET("/stats/system", api.ConvertHttpRouterToGin(GetSystemStats))
	router.GET("/stats/users", api.ConvertHttpRouterToGin(GetUserStats))
}

func GetRetrievalStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetStorageStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetSystemStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetUserStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func GetPublicStats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
