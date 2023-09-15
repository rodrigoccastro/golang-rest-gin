package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func serviceClientFind(c *gin.Context) gin.H {
	queryParam, _ := c.GetQuery("id");
	db := getDb()
	defer db.Close()

	rows,_ := db.Query("SELECT * FROM clients where id=$1",queryParam)
	if rows.Next() {
		var item Client
		rows.Scan(&item.id, &item.name, &item.email, &item.phone, &item.address, &item.created, &item.updated)
		return getResultDataItem(getDtoClient(item))

	}
	return getResultDataItem(nil)
}


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