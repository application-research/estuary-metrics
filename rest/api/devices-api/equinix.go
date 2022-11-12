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
	router.GET("/environment/equinix/usages", api.ConvertHttpRouterToGin(GetDeviceUsage))
	router.GET("/environment/equinix/list/usages", api.ConvertHttpRouterToGin(GetDeviceUsage))
	router.GET("/environment/equinix/info", api.ConvertHttpRouterToGin(GetDeviceInfo))
	router.GET("/environment/equinix/billing", api.ConvertHttpRouterToGin(GetBillingUsage))
}

func GetDeviceInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := r.FormValue("uuid")

	createdBefore := r.FormValue("createdBefore")
	createdAfter := r.FormValue("createdAfter")

	fmt.Println("uuid: ", uuid)
	info, err := dao.Metrics.GetDeviceInfo(uuid, createdAfter, createdBefore)
	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}

func GetDeviceUsage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := r.FormValue("uuid")

	createdBefore := r.FormValue("createdBefore")
	createdAfter := r.FormValue("createdAfter")

	fmt.Println("uuid: ", uuid)
	info, err := dao.Metrics.GetDeviceUsage(uuid, createdAfter, createdBefore)
	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}

func GetBillingUsage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := r.FormValue("uuid")

	createdBefore := r.FormValue("createdBefore")
	createdAfter := r.FormValue("createdAfter")

	fmt.Println("uuid: ", uuid)
	info, err := dao.Metrics.GetDeviceInfo(uuid, createdAfter, createdBefore)

	//	get the hourly rate
	//info.Plan.Pricing.Hour

	//	compute

	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}
