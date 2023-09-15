package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func clientList(c *gin.Context) {
	c.JSON(http.StatusOK, serviceClientList())
}

func clientFind(c *gin.Context) {
	c.JSON(http.StatusOK, serviceClientFind(c))
}


func main() {
	fmt.Println("starting server...")
	router := gin.Default()

	router.GET("/api/client/list", clientList)
	router.GET("/api/client/find", clientFind)
	
	router.Run("localhost:8090")
}
