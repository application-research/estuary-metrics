package app

import (
	"fmt"
	_ "github.com/application-research/estuary-metrics/rest/docs"
	"github.com/application-research/estuary-metrics/rest/route"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func GinServer() (engine *gin.Engine, err error) {
	router := gin.Default()
	router.Use(CORSMiddleware())
	url := ginSwagger.URL("https://metrics-api.estuary.tech/swagger/doc.json") // The url pointing to API definition
	route.ConfigRouter(router)

	router.Static("/web", "./web")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.Run("0.0.0.0:3030")
	if err != nil {
		log.Fatalf("Error starting server, the error is '%v'", err)
	}

	return router, err
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("CORS")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
