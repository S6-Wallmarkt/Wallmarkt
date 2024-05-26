package models

// NewOrder model
type NewOrder struct {
	CustomerID string   `json:"customer_id" binding:"required"`
	Payed      bool     `json:"payed"`
	ProductIDs []string `json:"products" binding:"required"`
}
