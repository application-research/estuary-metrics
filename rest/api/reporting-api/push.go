package reportingapi

import (
	"net/http"

	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
)

func ConfigMetricsPushRouter(router gin.IRoutes) {
	router.POST("/reporting/push/log", api.ConvertHttpRouterToGin(AddPushToLog))
	router.POST("/reporting/filter/log", api.ConvertHttpRouterToGin(AddPushToLog))
}

func AddPushToLog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func PrefixFilterFromLogDb(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}
