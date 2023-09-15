package main

import (
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

func getDb() *sql.DB {
	connStr := "user=postgres password=postgres dbname=DbSalesSystem sslmode=disable"
	db, _ := sql.Open("postgres", connStr)
	return db
}

func getResultDataItem(data gin.H) gin.H {
	return gin.H{
			"time": time.Now(),
			"data": data,
		}
}

func getResultDataList(data []gin.H) gin.H {
	return gin.H{
			"time": time.Now(),
			"data": data,
		}
}