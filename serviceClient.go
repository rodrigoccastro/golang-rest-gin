package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

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


func getDtoClients(rows *sql.Rows) gin.H {
	var clients []gin.H
	for rows.Next() {
		var item Client
		rows.Scan(&item.id, &item.name, &item.email, &item.phone, &item.address, &item.created, &item.updated)
		clients = append(clients, getDtoClient(item))
	}
	return getResultDataList(clients)
}