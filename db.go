package main

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "host=localhost port=5432 user=postgres password=yourpassword dbname=lab02shop sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}
}
