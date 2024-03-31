package api

import (
	product "github.com/S6-Wallmarkt/Wallmarkt/services/product/internal"

	"github.com/gin-gonic/gin"
)

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Routes
	router.GET("/create", product.Create)

	return router
}
