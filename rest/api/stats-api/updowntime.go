package statsapi

import "github.com/gin-gonic/gin"

func ConfigHeartbeatRoute(router gin.IRoutes) {
	router.GET("/heartbeat/check")
	router.GET("/heartbeat/")
}
