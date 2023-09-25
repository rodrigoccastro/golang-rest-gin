package main

import "github.com/gin-gonic/gin"

func getDtoClient(item Client) gin.H {
	return gin.H{
			"id": item.id,
			"name": item.name,
			"email": item.email,
			"phone": item.phone,
			"address": item.address,
			"created": item.created,
			"updated": item.updated,
		}
}