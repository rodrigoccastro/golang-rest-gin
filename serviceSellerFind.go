package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func serviceSellerFind(id string) gin.H {
	rows, err := DB.Query("SELECT * FROM sellers where id=$1",id)

	if err != nil {
		return getResultError("serviceSellerFind", err)
	}
	
	if rows.Next() {
		var item Seller
		rows.Scan(&item.id, &item.name, &item.email, &item.phone, &item.address, &item.created, &item.updated)
		return getResultDataItem(getDtoSeller(item))
	}

	return getResultDataItem(nil)
}