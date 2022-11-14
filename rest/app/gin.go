package app

import (
	_ "github.com/application-research/estuary-metrics/rest/docs"
	"github.com/application-research/estuary-metrics/rest/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func GinServer() (engine *gin.Engine, err error) {
	router := gin.Default()
	url := ginSwagger.URL("http://localhost:3030/swagger/doc.json") // The url pointing to API definition
	group := router.Group("/api/v1")
	route.ConfigRouter(group)

	router.Static("/web", "./web")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run("0.0.0.0:3030")
	if err != nil {
		log.Fatalf("Error starting server, the error is '%v'", err)
	}

	return router, err
}
