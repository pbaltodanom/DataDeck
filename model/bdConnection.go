package model

import (	
	"database/sql"

	//"github.com/pbaltodanom/DataDeck/models"
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlite3Str = "sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open(sqlite3Str, filepath)
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}