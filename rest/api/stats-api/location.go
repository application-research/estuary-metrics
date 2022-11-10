package statsapi

import "github.com/gin-gonic/gin"

func ConfigLocationRoute(router gin.IRoutes) {

	router.GET("/location")
}
