package devicesapi

import (
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ConfigEquinixDevicesRouter(router gin.IRoutes) {
	router.GET("/equinix/usage", api.ConverHttpRouterToGin(GetUsage))
}

func GetUsage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := ps.ByName("uuid")
	createdBefore := ps.ByName("createdBefore")
	createdAfter := ps.ByName("createdAfter")

	info, err := dao.Metrics.GetDeviceInfo(uuid, createdBefore, createdAfter)
	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}
