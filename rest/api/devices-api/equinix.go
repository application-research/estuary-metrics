package devicesapi

import (
	"fmt"
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
)

func ConfigEquinixDevicesRouter(router gin.IRoutes) {
	router.GET("/environment/equinix/usages", api.ConvertHttpRouterToGin(GetDeviceUsage))
	router.POST("/environment/equinix/list/usages", api.ConvertHttpRouterToGin(GetDevicesUsages))
	router.GET("/environment/equinix/info", api.ConvertHttpRouterToGin(GetDeviceInfo))
	router.GET("/environment/equinix/billing", api.ConvertHttpRouterToGin(GetBillingUsage))
}

func ConfigAwsDevicesRouter(router gin.IRoutes) {

}

// GetDeviceInfo returns the device info
// @Summary Get device info
// @Description Get device info
// @Tags environment
// @Accept  json
// @Produce  json
// @Param uuid query string true "uuid"
// @Param createdBefore query string true "createdBefore"
// @Param createdAfter query string true "createdAfter"
// @Success 200 {object} core.DeviceInfo
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
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

type UuidGroup struct {
	Uuids         []string `json:"uuids"`
	CreatedBefore string   `json:"createdBefore"`
	CreatedAfter  string   `json:"createdAfter"`
}

// GetDevicesUsages returns the device info
// @Summary Get device usages
// @Description Get device usages
// @Tags environment
// @Accept  json
// @Produce  json
// @Param uuids query string true "uuids"
// @Param createdBefore query string true "createdBefore"
// @Param createdAfter query string true "createdAfter"
// @Success 200 {object} []core.DeviceUsage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
func GetDevicesUsages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	var uuidGroup UuidGroup
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal(resBody, &uuidGroup)
	createdBefore := uuidGroup.CreatedBefore
	createdAfter := uuidGroup.CreatedAfter

	fmt.Println("uuid: ", uuidGroup)
	info, err := dao.Metrics.GetAllDeviceUsages(uuidGroup.Uuids, createdAfter, createdBefore)
	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}

// GetDeviceUsage returns the device usage
// @Summary Get device usage
// @Description Get device usage
// @Tags environment
// @Accept  json
// @Produce  json
// @Param uuid query string true "uuid"
// @Param createdBefore query string true "createdBefore"
// @Param createdAfter query string true "createdAfter"
// @Success 200 {object} core.DeviceUsage
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
func GetDeviceUsage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := r.FormValue("uuid")

	createdBefore := r.FormValue("createdBefore")
	createdAfter := r.FormValue("createdAfter")

	info, err := dao.Metrics.GetDeviceUsage(uuid, createdAfter, createdBefore)
	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}

// GetBillingUsage returns the device usage
// @Summary Get device usage
// @Description Get device usage
// @Tags environment
// @Accept  json
// @Produce  json
// @Param uuid query string true "uuid"
// @Param createdBefore query string true "createdBefore"
// @Param createdAfter query string true "createdAfter"
// @Success 200 {object} core.DeviceInfo
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
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
