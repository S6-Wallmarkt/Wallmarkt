package models

// Order model
type Order struct {
	ID         string   `json:"id" bson:"_id, omitempty"`
	CustomerID string   `json:"customer_id" binding:"required" bson:"customer_id"`
	Payed      bool     `json:"payed" binding:"required" bson:"payed"`
	Shipped    bool     `json:"shipped" bson:"shipped"`
	Products   []string `json:"products" binding:"required" bson:"products"`
}
