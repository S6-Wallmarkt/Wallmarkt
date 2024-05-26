package repository

import (
	"context"
	"time"
)

type OrderRepository struct{}

func (p OrderRepository) AddOrder() {
	// Set up context and close when done
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
}
