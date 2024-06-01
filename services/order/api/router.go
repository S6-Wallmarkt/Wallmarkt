package api

import (
	"fmt"
	"net/http"

	orderLogic "github.com/S6-Wallmarkt/Wallmarkt/services/order/internal"
	"github.com/S6-Wallmarkt/Wallmarkt/services/order/models"
	"github.com/gin-gonic/gin"
)

const internalServerErrorMessage string = "Something went wrong, try again later"

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// POST: Create Order
	router.POST("/add", func(c *gin.Context) {
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

		id, err := orderLogic.Create(order)

		if err != nil {
			// Return internal error if failed
			c.JSON(http.StatusBadRequest, gin.H{"error": internalServerErrorMessage})
			return
		}

		// Return the created product
		c.JSON(http.StatusCreated, id)
	})

	// GET: Get by customer
	router.GET("/getbycustomer/:id", func(c *gin.Context) {
		customerID := c.Param("id")

		orders, err := orderLogic.GetByCustomer(customerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
			return
		}

		// If no orders found
		if len(orders) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
			return
		}

		c.JSON(http.StatusOK, orders)
	})

	// GET: Get order by ID
	router.GET("/getbyid/:id", func(c *gin.Context) {
		orderID := c.Param("id")

		order, err := orderLogic.GetByID(orderID)
		if err != nil {
			if err.Error() == "product not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
				return
			}
		}

		c.JSON(http.StatusOK, order)
	})

	// GET: Get all orders
	router.GET("/getall", func(c *gin.Context) {
		orders, err := orderLogic.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
			return
		}

		// If no orders found
		if len(orders) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No orders found"})
			return
		}

		c.JSON(http.StatusOK, orders)
	})

	// DELETE: Delete order
	router.DELETE("/delete/:id", func(c *gin.Context) {
		orderID := c.Param("id")

		err := orderLogic.Delete(orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
		}
		c.JSON(http.StatusOK, fmt.Sprintf("Deleted order: %v", orderID))
	})
	return router
}
