package internal

import (
	"testing"

	"github.com/S6-Wallmarkt/Wallmarkt/services/product/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// Test the GetAll function
	products := GetAll()
	assert.NotNil(t, products)
	assert.NotEmpty(t, products)
	assert.Equal(t, 5, len(products))
}

func TestGetByID(t *testing.T) {
	// Test with a valid ID
	product, err := GetByID(1) // Assuming 1 is a valid ID in your products slice
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, 1, product.ID)

	// Test with an invalid ID
	product, err = GetByID(999) // Assuming 999 is not a valid ID in your products slice
	assert.Error(t, err)
	assert.Empty(t, product)
}

func TestGetByType(t *testing.T) {
	// Test with a valid type
	products := GetByType("Lighting")
	assert.NotEmpty(t, products)
	assert.Equal(t, 3, len(products))

	products = GetByType("Non-existent type")
	assert.Empty(t, products)
}

func TestCreate(t *testing.T) {
	// Test with a valid product
	product := models.Product{
		ID:          6,
		Name:        "Test",
		Description: "Test",
		Price:       0,
		Color:       "Test",
		Types:       []string{"Test"},
	}
	addedProduct, err := Create(product)
	assert.NoError(t, err)
	assert.Equal(t, product, addedProduct)

	// Test with an invalid product
	product = models.Product{
		ID:          3,
		Name:        "",
		Description: "Test",
		Price:       0,
		Color:       "Test",
		Types:       []string{"Test"},
	}
	addedProduct, err = Create(product)
	assert.Error(t, err)
	assert.Empty(t, addedProduct)
	assert.Equal(t, "product already exists", err.Error())
}
