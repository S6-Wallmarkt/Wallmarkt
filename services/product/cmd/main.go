package main

import (
	router "github.com/S6-Wallmarkt/Wallmarkt/services/product/api"
)

func main() {
	router := router.SetupRouter()
	err := router.Run(":8081")
	if err != nil {
		panic(err)
	}
}
