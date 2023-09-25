package main

import "database/sql"

var DB *sql.DB

func startConnection() {
	connStr := "user=postgres password=postgres dbname=DbSalesSystem sslmode=disable"
	db, _ := sql.Open("postgres", connStr)
	DB = db
}
