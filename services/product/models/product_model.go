package models

// Product struct
type Product struct {
	ID             string   `json:"id"`
	Name           string   `json:"name" binding:"required"`
	Description    string   `json:"description" binding:"required"`
	Price          uint     `json:"price" binding:"required"`
	Color          string   `json:"color"`
	Types          []string `json:"types"`
	Stock          uint16   `json:"stock"`
	AvailableStock uint16   `json:"available_stock"`
}
