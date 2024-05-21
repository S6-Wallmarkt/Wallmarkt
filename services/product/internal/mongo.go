package internal

import (
	"context"

	log "github.com/sirupsen/logrus"

	"errors"
	"os"

	"github.com/S6-Wallmarkt/Wallmarkt/services/product/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func InitMongoDB(uri string) {
	clientOptions := options.Client().ApplyURI(uri)
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = Client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Connected to MongoDB Instance")
}

// Assume you have a collection named "products"
var Collection *mongo.Collection

func InitCollections() {
	collection := Client.Database(os.Getenv("MONGO_DATABASE")).Collection(os.Getenv("MONGO_COLLECTION"))
	Collection = collection
}

// Get function to get all products
func GetAllProducts() ([]models.Product, error) {
	// Empty filter to find all documents
	cursor, err := Collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Decode and return all documents
	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Error(err)
		return nil, err
	}

	return products, nil
}

// Get function to get product by id
func GetProductByID(id string) (models.Product, error) {
	product := models.Product{}
	objectID, errId := primitive.ObjectIDFromHex(id)
	if errId != nil {
		log.Error(errId)
		return models.Product{}, errId
	}

	// Set id
	product.ID = id

	err := Collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Product{}, errors.New("product not found")
		}
		log.Error(err)
		return models.Product{}, err
	}
	return product, nil
}

// Get function to get all products that have the given type
func GetProductsWithType(_type string) ([]models.Product, error) {
	cursor, err := Collection.Find(context.TODO(), bson.M{"types": _type})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// Decode and return all documents
	var products []models.Product
	if err = cursor.All(context.TODO(), &products); err != nil {
		log.Error(err)
		return nil, err
	}

	return products, nil
}

// Create function to add products
func CreateProduct(product models.Product) (primitive.ObjectID, error) {
	// Generate a new ObjectID for the product
	id := primitive.NewObjectID()

	// Prepare the product document with the generated ID
	productDocument := bson.M{
		"_id":         id,
		"name":        product.Name,
		"description": product.Description,
		"price":       product.Price,
		"color":       product.Color,
		"types":       product.Types,
	}

	// Insert the product document into the collection
	result, err := Collection.InsertOne(context.TODO(), productDocument)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	// Return the inserted document's ID
	return result.InsertedID.(primitive.ObjectID), nil
}
