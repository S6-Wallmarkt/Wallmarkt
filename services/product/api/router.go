package api

import (
	"net/http"
	"strconv"

	productLogic "github.com/S6-Wallmarkt/Wallmarkt/services/product/internal"
	"github.com/S6-Wallmarkt/Wallmarkt/services/product/models"
	"github.com/gin-gonic/gin"
)

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// GET: All products
	router.GET("/getall", func(c *gin.Context) {
		products := productLogic.GetAll()
		c.JSON(http.StatusOK, products)
	})

	// GET: Product by ID
	router.GET("/getbyid/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Convert id to an integer if fails return error
		productId, err := strconv.Atoi(id)
		if err != nil {
			// Handle the case where id is not a valid integer
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		// Find the product with the given id and return it
		product, err := productLogic.GetByID(productId)
		if err != nil {
			// If no product found with given id return error
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}

		// Return the product
		c.JSON(http.StatusOK, product)
	})

	// GET: Products by type
	router.GET("/getbytype/:type", func(c *gin.Context) {
		// Check for type param, if not found return an error
		_type := c.Param("type")

		// Find the products with the given type and return them
		products := productLogic.GetByType(_type)

		// Check if products is empty, if so return an error
		if len(products) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Products not found"})
			return
		}

		// Return the products
		c.JSON(http.StatusOK, products)
	})

	// POST: Add product
	router.POST("/add", func(c *gin.Context) {
		var product models.Product

		// Bind the JSON data from the request body to the product struct
		if err := c.ShouldBindJSON(&product); err != nil {
			// Handle the case where JSON binding fails
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
			return
		}

		// Create the product
		addedProduct, err := productLogic.Create(product)
		if err != nil {
			// Handle the case where product creation fails
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Return the created product
		c.JSON(http.StatusCreated, addedProduct)
	})

	return router
}
