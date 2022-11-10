package reportingapi

import (
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ConfigMetricsPushRouter(router gin.IRoutes) {
	router.POST("/push/log", api.ConverHttpRouterToGin(AddPushToLog))
}

func AddPushToLog(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
