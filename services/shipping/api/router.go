package api

import (
	"fmt"
	"net/http"

	shippingLogic "github.com/S6-Wallmarkt/Wallmarkt/services/shipping/internal"
	"github.com/S6-Wallmarkt/Wallmarkt/services/shipping/models"
	"github.com/gin-gonic/gin"
)

const internalServerErrorMessage string = "Something went wrong, try again later"

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// POST: Create Shipment
	router.POST("/add", func(c *gin.Context) {
		var shipment models.Shipment

		// Bind the JSON data from the request body to the shipment struct
		if err := c.ShouldBindJSON(&shipment); err != nil {
			// Handle the case where JSON binding fails
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		id, err := shippingLogic.Add(shipment)

		if err != nil {
			// Return internal error if failed
			c.JSON(http.StatusBadRequest, gin.H{"error": internalServerErrorMessage})
			return
		}

		// Return the created product
		c.JSON(http.StatusCreated, id)
	})

	// GET: Get shipment by ID
	router.GET("/getbyid/:id", func(c *gin.Context) {
		shipmentID := c.Param("id")

		shipment, err := shippingLogic.GetByID(shipmentID)
		if err != nil {
			if err.Error() == "shipment not found" {
				c.JSON(http.StatusNotFound, gin.H{"error": "shipment not found"})
				return
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
				return
			}
		}

		c.JSON(http.StatusOK, shipment)
	})

	// GET: Get all shipments
	router.GET("/getall", func(c *gin.Context) {
		shipments, err := shippingLogic.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
			return
		}

		// If no orders found
		if len(shipments) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No shipments found"})
			return
		}

		c.JSON(http.StatusOK, shipments)
	})

	// GET: Get all unsend shipments
	router.GET("/getallunsend", func(c *gin.Context) {
		shipments, err := shippingLogic.GetAllUnsend()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
			return
		}

		// If no orders found
		if len(shipments) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No shipments found"})
			return
		}

		c.JSON(http.StatusOK, shipments)
	})

	// DELETE: Delete shipment
	router.DELETE("/delete/:id", func(c *gin.Context) {
		shipmentID := c.Param("id")

		err := shippingLogic.Delete(shipmentID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": internalServerErrorMessage})
		}
		c.JSON(http.StatusOK, fmt.Sprintf("Deleted shipment: %v", shipmentID))
	})
	return router
}
