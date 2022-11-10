package devicesapi

import (
	"fmt"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ConfigEquinixDevicesRouter(router gin.IRoutes) {
	router.GET("/environment/equinix/usage", api.ConvertHttpRouterToGin(GetUsage))
}

func GetUsage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := r.FormValue("uuid")

	createdBefore := r.FormValue("createdBefore")
	createdAfter := r.FormValue("createdAfter")

	fmt.Println("uuid: ", uuid)
	info, err := dao.Metrics.GetDeviceInfo(uuid, createdBefore, createdAfter)
	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}
