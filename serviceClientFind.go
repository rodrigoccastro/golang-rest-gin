package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func serviceClientFind(id string) gin.H {
	rows, err := DB.Query("SELECT * FROM clients where id=$1",id)

	if err != nil {
		return getResultError(err)
	}
	
	if rows.Next() {
		var item Client
		rows.Scan(&item.id, &item.name, &item.email, &item.phone, &item.address, &item.created, &item.updated)
		return getResultDataItem(getDtoClient(item))
	}

	return getResultDataItem(nil)
}