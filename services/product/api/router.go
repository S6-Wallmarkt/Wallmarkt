package api

import (
	product "github.com/S6-Wallmarkt/Wallmarkt/services/product/internal"

	"github.com/gin-gonic/gin"
)

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Define routes
	router.GET("/getall", product.GetAll)
	router.GET("/getbyid/:id", product.GetByID)
	router.GET("/getbytype/:type", product.GetByType)

	router.POST("/add", product.Create)

	return router
}
