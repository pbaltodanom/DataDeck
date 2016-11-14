package api

//---------------- IMPORTS ---------------- 
import (	
	"database/sql"

	//"github.com/pbaltodanom/DataDeck/models"
	_ "github.com/mattn/go-sqlite3"
)

const (
	sqlite3Str = "sqlite3"
)

//filepath = path where is located the  db
func InitDB(filepath string) *sql.DB {
	db, err := sql.Open(sqlite3Str, filepath)
	//Opens the database specified by the driver name (sqlite3Str) and a driver-specific data source name (the path - filepath variable), 
	if err != nil { panic(err) }
	if db == nil { panic("db nil") }
	return db
}