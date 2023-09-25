package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func controllerSellerList(c *gin.Context) {
	c.JSON(http.StatusOK, serviceSellerList())
}

func controllerSellerFind(c *gin.Context) {
	c.JSON(http.StatusOK, serviceSellerFind(c.Param("id")))
}