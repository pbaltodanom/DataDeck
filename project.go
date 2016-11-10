package main

import (
	"log"
	"database/sql"

	"github.com/pbaltodanom/DataDeck/models"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}


func main() {
	const dbpath = "https://s3.amazonaws.com/bv-challenge/jrdd.db"

	db := InitDB(dbpath)
	defer db.Close()

	me := models.Songs{Artist: "Santana", Name:  "Smooth", Genre: "Pop", Length: 167}

	log.Printf("SELECT")
	selectedSongs := models.Songs{}
	if err := models.Query(db, me.Artist, me.Name, me.Genre, me.Length, &selectedSongs); err != nil {
		log.Fatalf("Error selecting person from the DB (%s)", err)
	}
	log.Printf("Selected %+v from the DB", selectedSongs)
}