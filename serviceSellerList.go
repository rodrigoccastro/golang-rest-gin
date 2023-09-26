package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func serviceSellerList() gin.H {
	rows, err := DB.Query("SELECT * FROM sellers order by id asc")

	if err != nil {
		return getResultError("serviceSellerList", err)
	}
	
	return getDtoSellers(rows)
}

func getDtoSellers(rows *sql.Rows) gin.H {
	var sellers []gin.H
	for rows.Next() {
		var item Seller
		rows.Scan(&item.id, &item.name, &item.email, &item.phone, &item.address, &item.created, &item.updated)
		sellers = append(sellers, getDtoSeller(item))
	}
	return getResultDataList(sellers)
}

