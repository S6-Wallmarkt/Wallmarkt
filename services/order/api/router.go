package api

import (
	"net/http"

	"github.com/S6-Wallmarkt/Wallmarkt/services/order/models"
	"github.com/gin-gonic/gin"
)

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// POST: Create Order
	router.POST("/create", func(c *gin.Context) {
		var order models.Order

		// Bind the JSON data from the request body to the order struct
		if err := c.ShouldBindJSON(&order); err != nil {
			// Handle the case where JSON binding fails
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		if order.Products == nil || len(order.Products) == 0 {
			// Handle the case where no products are provided
			c.JSON(http.StatusBadRequest, gin.H{"error": "No products provided"})
			return
		}

		// Return the created product
		c.JSON(http.StatusCreated, order)
	})
	return router
}
