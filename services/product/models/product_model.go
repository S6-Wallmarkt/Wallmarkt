package models

// Product struct
type Product struct {
	ID             string   `json:"id" bson:"_id,omitempty"`
	Name           string   `json:"name" binding:"required" bson:"name"`
	Description    string   `json:"description" binding:"required" bson:"description"`
	Price          uint     `json:"price" binding:"required" bson:"price"`
	Color          string   `json:"color" bson:"color"`
	Types          []string `json:"types" bson:"types"`
	Stock          uint16   `json:"stock" bson:"stock"`
	AvailableStock uint16   `json:"available_stock" bson:"available_stock"`
}
