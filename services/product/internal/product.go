package internal

import (
	"errors"

	"github.com/S6-Wallmarkt/Wallmarkt/services/product/models"
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
func GetAll() []models.Product {
	return products
}

// Get function to get product by id
func GetByID(id int) (models.Product, error) {
	for _, product := range products {
		if product.ID == id {
			return product, nil
		}
	}
	return models.Product{}, errors.New("product not found")
}

// Get function to get products by type
func GetByType(_type string) []models.Product {
	filteredProducts := []models.Product{}

	// Add product to filteredProducts if it has the given type
	for _, product := range products {
		for _, productType := range product.Types {
			if productType == _type {
				filteredProducts = append(filteredProducts, product)
				break
			}
		}
	}

	return filteredProducts
}

// Post function to create new product
func Create(product models.Product) (models.Product, error) {
	// Check if the product already exists
	for _, p := range products {
		if p.ID == product.ID {
			return models.Product{}, errors.New("product already exists")
		}
	}

	// Add the new product to the products slice
	products = append(products, product)

	return product, nil
}
