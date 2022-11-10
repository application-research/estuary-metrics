package statsapi

import "github.com/gin-gonic/gin"

func ConfigRankingRoute(router gin.IRoutes) {

	router.GET("/rank/miners/:top")
	router.GET("/rank/users/:top")
	router.GET("/rank/users/collection/:top")

}
