package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func serviceClientList() gin.H {
	rows, err := DB.Query("SELECT * FROM clients order by id asc")

	if err != nil {
		return getResultError("serviceClientList", err)
	}
	
	return getDtoClients(rows)
}
