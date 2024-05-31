package repository

import (
	"context"
	"errors"
	"time"

	"github.com/S6-Wallmarkt/Wallmarkt/services/shipping/configs"
	"github.com/S6-Wallmarkt/Wallmarkt/services/shipping/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

type ShipmentRepository struct{}

func (p ShipmentRepository) AddShipment(shipment models.Shipment) (string, error) {
	// Generate a new ObjectID for the product
	id := primitive.NewObjectID()

	// Prepare the product document with the generated ID
	shipmentDocument := bson.M{
		"_id":       id,
		"order":     shipment.Order,
		"timestamp": time.Now(),
		"send":      shipment.Send,
		"packer":    shipment.Packer,
	}
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Add order
	result, err := configs.Collection.InsertOne(ctx, shipmentDocument)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Return the inserted document's ID
	return result.InsertedID.(primitive.ObjectID).String(), nil
}

func (p ShipmentRepository) GetShipmentByID(shipmentID string) (models.Shipment, error) {
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var shipment models.Shipment

	// Turn into objectID
	objectID, errId := primitive.ObjectIDFromHex(shipmentID)
	if errId != nil {
		log.Error(errId)
		return models.Shipment{}, errId
	}

	// Find shipment by ID
	err := configs.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&shipment)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Shipment{}, errors.New("shipment not found")
		}
		log.Error(err)
		return models.Shipment{}, err
	}
	return shipment, nil
}

func (p ShipmentRepository) GetAllShipments() ([]models.Shipment, error) {
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all shipments
	cursor, err := configs.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Bind items to model
	var shipments []models.Shipment
	if err = cursor.All(ctx, &shipments); err != nil {
		log.Error(err)
		return nil, err
	}

	return shipments, nil
}

func (p ShipmentRepository) GetAllUnsendShipments() ([]models.Shipment, error) {
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := configs.Collection.Find(ctx, bson.M{"send": false})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Decode and return all documents
	var shipments []models.Shipment
	if err = cursor.All(context.TODO(), &shipments); err != nil {
		log.Error(err)
		return nil, err
	}

	return shipments, nil
}
