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

	router.GET("/api/client/list", controllerClientList)
	router.GET("/api/client/find/:id", controllerClientFind)
	
	router.GET("/api/seller/list", controllerSellerList)
	router.GET("/api/seller/find/:id", controllerSellerFind)

	router.Run("localhost:8090")
}