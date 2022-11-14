package statsapi

import (
	"github.com/application-research/estuary-metrics/core/dao"
	"github.com/application-research/estuary-metrics/rest/api"
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func ConfigLocationRoute(router gin.IRoutes) {
	router.GET("/location/shuttle/:uuid", api.ConvertHttpRouterToGin(GetShuttleLocation))
	router.GET("/location/shuttles", api.ConvertHttpRouterToGin(GetShuttleLocations))
}

type Location struct {
	Name    string `json:"name"`
	Address struct {
		ID          string `json:"id"`
		Address     string `json:"address"`
		Address2    string `json:"address2"`
		City        string `json:"city"`
		State       string `json:"state"`
		ZipCode     string `json:"zip_code"`
		Country     string `json:"country"`
		Coordinates struct {
			Latitude  string `json:"latitude"`
			Longitude string `json:"longitude"`
		} `json:"coordinates"`
	} `json:"address"`
	IPAddress struct {
		Network string `json:"network"`
		Address string `json:"address"`
		Gateway string `json:"gateway"`
	} `json:"ip_address"`
}

// GetShuttleLocation returns the device info
// @Summary Get device info
// @Description Get device info
// @Tags environment
// @Accept  json
// @Produce  json
// @Param uuid path string true "uuid"
// @Success 200 {object} Location
// @Failure 400 {object} api.HTTPError
// @Failure 500 {object} api.HTTPError
// @Router /location/shuttle/{uuid} [get]
func GetShuttleLocation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := api.InitializeContext(r)
	var locationInfo Location
	uuid := ps.ByName("uuid")

	deviceInfo, err := dao.Metrics.GetDeviceInfo(uuid)
	if deviceInfo == nil {
		api.ReturnError(ctx, w, r, err)
		return
	}

	locationInfo.IPAddress.Address = deviceInfo.IPAddresses[0].Address
	locationInfo.IPAddress.Gateway = deviceInfo.IPAddresses[0].Gateway
	locationInfo.IPAddress.Network = deviceInfo.IPAddresses[0].Network
	locationInfo.Address = deviceInfo.Facility.Address
	locationInfo.Name = deviceInfo.Facility.Name

	api.WriteJSON(ctx, w, locationInfo)
}

func GetShuttleLocations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}
