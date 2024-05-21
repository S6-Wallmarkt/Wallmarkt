package main

import (
	"os"

	router "github.com/S6-Wallmarkt/Wallmarkt/services/product/api"
	internal "github.com/S6-Wallmarkt/Wallmarkt/services/product/internal"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {

	// Load the .env file
	// Check if the --development flag is provided
	if len(os.Args) > 1 && os.Args[1] == "--development" {
		log.SetLevel(log.DebugLevel)
		err := godotenv.Load()
		log.Print("Loading .env file")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Set up database connection
	mongodbUri := os.Getenv("MONGO_URI")
	internal.InitMongoDB(mongodbUri)
	internal.InitCollections()

	// Set up router
	router := router.SetupRouter()
	err := router.Run(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
}
