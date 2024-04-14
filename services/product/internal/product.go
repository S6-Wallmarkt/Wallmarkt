package internal

import (
	"net/http"
	"strconv"

	"github.com/S6-Wallmarkt/Wallmarkt/services/product/models"
	"github.com/gin-gonic/gin"
)

// List of mock products
var products = []models.Product{
	{
		ID:          1,
		Name:        "Candle",
		Description: "Light up rooms with this candle",
		Price:       9.95,
		Color:       "White",
		Types:       []string{"Lighting", "Decorative"},
	},
	{
		ID:          2,
		Name:        "Wall-torch",
		Description: "Light up rooms with this wall-torch",
		Price:       25.50,
		Color:       "Steel",
		Types:       []string{"Lighting", "Decorative"},
	},
	{
		ID:          3,
		Name:        "Basket",
		Description: "Store items in a basket",
		Price:       5,
		Color:       "Wood",
		Types:       []string{"Decorative", "Storage"},
	},
	{
		ID:          4,
		Name:        "Star-lamp",
		Description: "Light up rooms with this wall mounted star lamp",
		Price:       14.95,
		Color:       "White",
		Types:       []string{"Lighting", "Decorative"},
	},
	{
		ID:          5,
		Name:        "Deer-mount",
		Description: "Creepy deer head mount for your wall",
		Price:       999.99,
		Color:       "Brown",
		Types:       []string{"Creepy", "Decorative"},
	},
}

// GetAll function to get all products
func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

// Get function to get product by id
func GetByID(c *gin.Context) {

	// Check for id param, if not found return an error
	id := c.Param("id")
	if id == "" {
		// Handle the case where id is an empty string
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID is required"})
		return
	}

	// Convert id to an integer if fails return error
	idInt, err := strconv.Atoi(id)
	if err != nil {
		// Handle the case where id is not a valid integer
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Find the product with the given id and return it
	for _, product := range products {
		if product.ID == idInt {
			c.JSON(http.StatusOK, product)
			return
		}
	}

	// If no product found with given id return error
	c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
}

// Get function to get products by type
func GetByType(c *gin.Context) {
	var filteredProducts []models.Product

	// Check for type param, if not found return an error
	_type := c.Param("type")
	if _type == "" {
		// Handle the case where type is an empty string
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type is required"})
		return
	}

	// Add product to filteredProducts if it has the given type
	for _, product := range products {
		for _, productType := range product.Types {
			if productType == c.Param("type") {
				filteredProducts = append(filteredProducts, product)
				break
			}
		}
	}

	// If no products found with given type return error
	if len(filteredProducts) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Products not found"})
		return
	}

	// Return the filtered products
	c.JSON(http.StatusOK, filteredProducts)
}

// Post function to create new product
func Create(c *gin.Context) {
	var product models.Product

	// Bind the JSON data from the request body to the product struct
	if err := c.ShouldBindJSON(&product); err != nil {
		// Handle the case where JSON binding fails
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	// Add the new product to the products slice
	products = append(products, product)

	// Return the created product
	c.JSON(http.StatusCreated, product)
}
