package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Static("/web", "./web")

	// TODO: Parameterize
	router.Run("localhost:3031")

}
