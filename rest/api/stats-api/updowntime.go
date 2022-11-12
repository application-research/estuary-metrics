package statsapi

import "github.com/gin-gonic/gin"

//Downtime
//Performance
func ConfigHeartbeatRoute(router gin.IRoutes) {
	router.GET("/heartbeat/check")
	router.GET("/heartbeat/")
}
