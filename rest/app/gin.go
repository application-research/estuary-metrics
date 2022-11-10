package app

import (
	"github.com/application-research/estuary-metrics/rest/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func GinServer() (err error) {
	url := ginSwagger.URL("http://localhost:3030/swagger/doc.json") // The url pointing to API definition

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Group("/api/v0")
	route.ConfigRouter(router)

	router.Static("/web", "./web")

	// TODO: Parameterize
	router.Run("localhost:3030")
	if err != nil {
		log.Fatalf("Error starting server, the error is '%v'", err)
	}

	return
}
