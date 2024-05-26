package internal

import (
	"github.com/S6-Wallmarkt/Wallmarkt/services/order/models"
	"github.com/S6-Wallmarkt/Wallmarkt/services/order/repository"
)

var OrderRepository repository.OrderRepository

func CreateOrder(order models.NewOrder) (string, error) {
	OrderRepository.AddOrder()
	return "", nil
}
