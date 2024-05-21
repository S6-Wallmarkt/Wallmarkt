package api

import (
	"fmt"
	"log"
	"net/http"

	productLogic "github.com/S6-Wallmarkt/Wallmarkt/services/product/internal"
	"github.com/S6-Wallmarkt/Wallmarkt/services/product/models"
	"github.com/gin-gonic/gin"
)

const InternalServerErrorMessage string = "Something went wrong, try again later"

// Function to setup the router and its routes
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// GET: All products
	router.GET("/getall", func(c *gin.Context) {
		products, err := productLogic.GetAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMessage})
		}
		c.JSON(http.StatusOK, products)
	})

	// GET: Product by ID
	router.GET("/getbyid/:id", func(c *gin.Context) {
		id := c.Param("id")

		// Find the product with the given id and return it
		product, err := productLogic.GetByID(id)
		if err != nil {
			if err.Error() == "product not found" {
				// If no product found with given id return error
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMessage})
			}
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
		products, err := productLogic.GetByType(_type)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMessage})
			return
		}
		// Check if products is empty, if so return an error
		if len(products) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Products not found by type: %v", _type)})
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
		newId, err := productLogic.Create(product)
		if err != nil {
			// Handle the case where product creation fails
			log.Fatal(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": InternalServerErrorMessage})
			return
		}

		// Return the created product
		c.JSON(http.StatusCreated, newId)
	})

	return router
}
