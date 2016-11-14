package api

import (
	"log"
	"encoding/json"
	"net/http"	

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

//PATH where is located the  db
const dbpath = "jrdd.db"

func GetArtist(w http.ResponseWriter, req *http.Request) {
	//Initialize the database
	db := InitDB(dbpath)
	defer db.Close()

	//// Vars returns the route variables for the current request, if any.
	params := mux.Vars(req)
	//A slice or array of type SONGS, this is where is goint to be all data retrieve from the db
	selectedSongs := make([]Songs, 0)

	//Call the query function
	if err := QueryGetArtist(db, params["artist"], &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}
	
	//For each row that was returned by the query
	for _, bk := range selectedSongs {
		//Display the data as a JSON values to an output stream.
		json.NewEncoder(w).Encode(bk)
  	}
}

func GetSong(w http.ResponseWriter, req *http.Request) {
	db := InitDB(dbpath)
	defer db.Close()

	params := mux.Vars(req)
	selectedSongs := make([]Songs, 0)

	if err := QueryGetSong(db, params["song"], &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}
	
	for _, bk := range selectedSongs {
		json.NewEncoder(w).Encode(bk)
  	}
}

func GetGenre(w http.ResponseWriter, req *http.Request) {
	db := InitDB(dbpath)
	defer db.Close()

	params := mux.Vars(req)
	selectedSongs := make([]Songs, 0)

	if err := QueryGetGenre(db, params["genre"], &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}

	for _, bk := range selectedSongs {
		json.NewEncoder(w).Encode(bk)
  	}
}


//Extra #1: Make a function in the API return songs by length, which we would like to search by passing a minimum and maximum length.
func GetSongs(w http.ResponseWriter, req *http.Request) {
	db := InitDB(dbpath)
	defer db.Close()

	params := mux.Vars(req)
	selectedSongs := make([]Songs, 0)

	if err := QueryGetSongs(db, params["min"], params["max"], &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}

	for _, bk := range selectedSongs {
		json.NewEncoder(w).Encode(bk)
  	}
}

//Extra #2: Make a function in the API return a list of the genres, and the number of songs and the total length of all the songs by genre
func GetGenres(w http.ResponseWriter, req *http.Request) {
	db := InitDB(dbpath)
	defer db.Close()

	selectedSongs := make([]Songs, 0)

	if err := QueryGetGenres(db, &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}

	for _, bk := range selectedSongs {
		json.NewEncoder(w).Encode(bk)
  	}
}

func Handlers() *mux.Router{
    //this is the object that will route our requests to the right handler function.
	router := mux.NewRouter()	

	//GetArtist function will handle GET /artist requests.
	router.HandleFunc("/artist/{artist}", GetArtist).Methods("GET")

	//GetSong function will handle GET /song requests.
	router.HandleFunc("/song/{song}", GetSong).Methods("GET")

	//GetGenre function will handle GET /genre requests.
	router.HandleFunc("/genre/{genre}", GetGenre).Methods("GET")

	//GetGenres function will handle GET /genres requests.
	router.HandleFunc("/genres", GetGenres).Methods("GET")

	//GetSongs function will handle GET /songs requests. Need two parameters the min and max lenght
	router.HandleFunc("/songs/min/{min}/max/{max}", GetSongs).Methods("GET")

    return router
}