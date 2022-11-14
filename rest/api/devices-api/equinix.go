package devicesapi

import (
	"fmt"
	"github.com/application-research/estuary-metrics/core"
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
// @Router /environment/equinix/info [get]
func GetDeviceInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := r.FormValue("uuid")
	fmt.Println("uuid: ", uuid)
	info, err := dao.Metrics.GetDeviceInfo(uuid)
	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}

//Backup: e4d0efb1-1b5b-4aaf-a6ed-37c4a6cc2c6f
//Shuttle-8: 8ceea3cd-7608-4428-8d6b-99f2acc80ce3
//Shuttle-7: 3c924716-f30e-4afd-a073-98204e4a96a7
//Shuttle-6: 266fbb9d-56a1-4dea-9b99-9f28054c5522
//
//Shuttle-2: ed16760d-ec36-4d71-b46f-378428c1d774
//Shuttle-1: 60352064-7b2c-4597-baf6-9df128e9242b
//
//Upload: a8e5d22b-13ef-4dc9-adcf-a3b2bb4a8863
//API: 766557e4-1c14-4bef-a5b2-d974bbb2d848

//
//func GetAllEquinixDeviceInfoUsage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	ctx := api.InitializeContext(r)
//	var uuidGroup UuidGroup
//	resBody, err := ioutil.ReadAll(r.Body)
//	if err != nil {
//
//		return
//	}
//	json.Unmarshal(resBody, &uuidGroup)
//	createdBefore := uuidGroup.CreatedBefore
//	createdAfter := uuidGroup.CreatedAfter
//
//	fmt.Println("uuid: ", uuidGroup)
//	info, err := dao.Metrics.GetAllDeviceUsages(uuidGroup.Uuids, createdAfter, createdBefore)
//	if err != nil {
//		return
//	}
//
//	api.WriteJSON(ctx, w, info)
//}

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
// @Router /environment/equinix/list/usages [post]
func GetDevicesUsages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	var uuidGroup core.UuidGroup
	resBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	json.Unmarshal(resBody, &uuidGroup)
	createdBefore := uuidGroup.CreatedBefore
	createdAfter := uuidGroup.CreatedAfter

	fmt.Println("uuid: ", uuidGroup)
	info, err := dao.Metrics.GetAllDeviceUsages(uuidGroup, createdAfter, createdBefore)
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
// @Router /environment/equinix/usages [get]
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
// @Router /environment/equinix/billing [get]
func GetBillingUsage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	uuid := r.FormValue("uuid")

	fmt.Println("uuid: ", uuid)
	info, err := dao.Metrics.GetDeviceInfo(uuid)

	//	get the hourly rate
	//info.Plan.Pricing.Hour

	//	compute

	if err != nil {
		return
	}

	api.WriteJSON(ctx, w, info)
}
