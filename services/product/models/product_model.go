package models

// Product struct
type Product struct {
	ID          int      `json:"id" binding:"required"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Price       float32  `json:"price" binding:"required"`
	Color       string   `json:"color"`
	Types       []string `json:"types"`
}
