package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func controllerClientList(c *gin.Context) {
	c.JSON(http.StatusOK, serviceClientList())
}

func controllerClientFind(c *gin.Context) {
	c.JSON(http.StatusOK, serviceClientFind(c.Param("id")))
}