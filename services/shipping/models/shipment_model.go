package models

import "time"

type Shipment struct {
	ID        string    `json:"id" bson:"_id, omitempty"`
	Order     string    `json:"order" binding:"required" bson:"order"`
	Timestamp time.Time `json:"timestamp" bson:"timestamp"`
	Send      bool      `json:"send" bson:"send"`
	Packer    string    `json:"packer" bson:"packer"`
}
