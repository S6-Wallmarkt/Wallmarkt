package api

import (
	"github.com/gin-gonic/gin"
)

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	return router
}
