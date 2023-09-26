package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/magiconair/properties"
)

var DB *sql.DB

func startConnection() {
	file, err := os.Open("database.config")
	if err != nil {
		fmt.Println("Error when read database.config")
		panic("Error when read database.config")
	}
	defer file.Close()

	p := properties.MustLoadFile("database.config", properties.UTF8)

	if err != nil {
		fmt.Println("Error when decode database.config")
		panic("Error when decode database.config")
	}

	connStr := "user="+p.MustGetString("user")+" password="+p.MustGetString("password")
	connStr = connStr + " dbname="+p.MustGetString("dbname")+" sslmode="+p.MustGetString("sslmode")

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error when connect database.json")
		panic("Error when connect database")
	}

	DB = db
}
