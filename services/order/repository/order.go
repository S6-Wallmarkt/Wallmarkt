package repository

import (
	"context"
	"errors"
	"time"

	"github.com/S6-Wallmarkt/Wallmarkt/services/order/configs"
	"github.com/S6-Wallmarkt/Wallmarkt/services/order/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	log "github.com/sirupsen/logrus"
)

type OrderRepository struct{}

func (p OrderRepository) AddOrder(order models.Order) (string, error) {
	// Generate a new ObjectID for the product
	id := primitive.NewObjectID()

	// Prepare the product document with the generated ID
	orderDocument := bson.M{
		"_id":         id,
		"customer_id": order.CustomerID,
		"payed":       order.Payed,
		"shipped":     order.Shipped,
		"products":    order.Products,
	}
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Add order
	result, err := configs.Collection.InsertOne(ctx, orderDocument)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Return the inserted document's ID
	return result.InsertedID.(primitive.ObjectID).String(), nil
}

func (p OrderRepository) GetOrdersByUser(customerID string) ([]models.Order, error) {
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a filter to find orders by customer_id
	filter := bson.M{"customer_id": customerID}

	// Find all orders by user
	cursor, err := configs.Collection.Find(ctx, filter)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Bind items to model
	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		log.Error(err)
		return nil, err
	}

	return orders, nil
}

func (p OrderRepository) GetOrderByID(orderID string) (models.Order, error) {
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order models.Order

	// Turn into objectID
	objectID, errId := primitive.ObjectIDFromHex(orderID)
	if errId != nil {
		log.Error(errId)
		return models.Order{}, errId
	}

	// Find order by ID
	err := configs.Collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&order)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Order{}, errors.New("product not found")
		}
		log.Error(err)
		return models.Order{}, err
	}
	return order, nil
}

func (p OrderRepository) GetAllOrders() ([]models.Order, error) {
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find all orders
	cursor, err := configs.Collection.Find(ctx, bson.D{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	// Bind items to model
	var orders []models.Order
	if err = cursor.All(ctx, &orders); err != nil {
		log.Error(err)
		return nil, err
	}

	return orders, nil
}

func (p OrderRepository) DeleteOrder(orderID string) error {
	// Set up context and close when done
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Turn into objectID
	objectID, errId := primitive.ObjectIDFromHex(orderID)
	if errId != nil {
		log.Error(errId)
		return errId
	}

	// Delete shipment
	_, err := configs.Collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
