package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB(uri string) *sql.DB {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
