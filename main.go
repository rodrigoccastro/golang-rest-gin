package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("starting server...")
	router := gin.Default()

	fmt.Println("starting connection...")
	startConnection();
	
	router.GET("/client", controllerClientList)
	router.GET("/client/:id", controllerClientFind)
	
	router.GET("/seller", controllerSellerList)
	router.GET("/seller/:id", controllerSellerFind)

	router.Run("localhost:8090")
}