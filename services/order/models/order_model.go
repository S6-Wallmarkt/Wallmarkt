package models

// Order model
type Order struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id" binding:"required"`
	Payed      bool      `json:"payed" binding:"required"`
	Products   []Product `json:"products" binding:"required"`
}

type Product struct {
	ID          uint     `json:"id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price       uint     `json:"price" binding:"required"`
	Color       string   `json:"color"`
	Types       []string `json:"types"`
}
