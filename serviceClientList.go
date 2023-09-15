package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func serviceClientList() gin.H {
	db := getDb()
	defer db.Close()

	rows, _ := db.Query("SELECT * FROM clients order by id asc")
	return getDtoClients(rows)
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

