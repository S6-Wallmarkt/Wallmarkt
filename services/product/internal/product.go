package internal

import (
	"log"

	"github.com/S6-Wallmarkt/Wallmarkt/services/product/models"
)

// GetAll function to get all products
func GetAll() ([]models.Product, error) {
	products, err := GetAllProducts()
	if err != nil {
		return products, err
	}

	return products, nil
}

// Get function to get product by id
func GetByID(id string) (models.Product, error) {
	product, err := GetProductByID(id)
	return product, err
}

// Get function to get products by type
func GetByType(_type string) ([]models.Product, error) {
	products, err := GetProductsWithType(_type)
	return products, err
}

// Post function to create new product
func Create(product models.Product) (string, error) {
	id, err := CreateProduct(product)
	if err != nil {
		log.Fatal(err)
	}

	return id.String(), nil
}
