package main

import (
	router "github.com/S6-Wallmarkt/Wallmarkt/services/order/api"
)

func main() {
	router := router.SetupRouter()
	err := router.Run(":8082")
	if err != nil {
		panic(err)
	}
}
