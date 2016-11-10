package main

import (
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3",
		"https://s3.amazonaws.com/bv-challenge/jrdd.db")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("se realizón conexión")
	}
	defer db.Close()
}