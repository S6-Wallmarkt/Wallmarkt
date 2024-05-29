package internal

import (
	"github.com/S6-Wallmarkt/Wallmarkt/services/order/models"
	"github.com/S6-Wallmarkt/Wallmarkt/services/order/repository"
)

var OrderRepository repository.OrderRepository

func Create(order models.Order) (string, error) {
	id, err := OrderRepository.AddOrder(order)
	if err != nil {
		return "", err
	}
	return id, nil
}

func GetByCustomer(customerID string) ([]models.Order, error) {
	orders, err := OrderRepository.GetOrdersByUser(customerID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func GetByID(orderID string) (models.Order, error) {
	order, err := OrderRepository.GetOrderByID(orderID)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func GetAll() ([]models.Order, error) {
	orders, err := OrderRepository.GetAllOrders()
	if err != nil {
		return []models.Order{}, err
	}
	return orders, nil
}

func Delete(orderID string) error {
	err := OrderRepository.DeleteOrder(orderID)
	if err != nil {
		return err
	}
	return nil
}
