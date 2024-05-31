package internal

import (
	"github.com/S6-Wallmarkt/Wallmarkt/services/shipping/models"
	"github.com/S6-Wallmarkt/Wallmarkt/services/shipping/repository"

	log "github.com/sirupsen/logrus"
)

var ShipmentRepository repository.ShipmentRepository

func Add(shipment models.Shipment) (string, error) {
	id, err := ShipmentRepository.AddShipment(shipment)
	if err != nil {
		log.Error(err)
		return "", err
	}
	return id, nil
}

func GetByID(shipmentID string) (models.Shipment, error) {
	shipment, err := ShipmentRepository.GetShipmentByID(shipmentID)
	if err != nil {
		log.Error(err)
		return models.Shipment{}, err
	}
	return shipment, nil
}

func GetAll() ([]models.Shipment, error) {
	shipments, err := ShipmentRepository.GetAllShipments()
	if err != nil {
		return []models.Shipment{}, err
	}
	return shipments, nil
}

func GetAllUnsend() ([]models.Shipment, error) {
	shipments, err := ShipmentRepository.GetAllUnsendShipments()
	if err != nil {
		return []models.Shipment{}, err
	}
	return shipments, nil
}
