package models

// Product struct
type Product struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	Color       string   `json:"color"`
	Types       []string `json:"types"`
}
