package main

import (
	"log"
	"encoding/json"
	"net/http"	

	model "DataDeck/model"	
	//"github.com/pbaltodanom/DataDeck/models"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

const dbpath = "jrdd.db"

func GetArtist(w http.ResponseWriter, req *http.Request) {
	db := model.InitDB(dbpath)
	defer db.Close()

	params := mux.Vars(req)
	selectedSongs := make([]model.Songs, 0)

	if err := model.QueryGetArtist(db, params["artist"], &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}
	
	for _, bk := range selectedSongs {
		json.NewEncoder(w).Encode(bk)
  	}
}

func GetSong(w http.ResponseWriter, req *http.Request) {
	db := model.InitDB(dbpath)
	defer db.Close()

	params := mux.Vars(req)
	selectedSongs := make([]model.Songs, 0)

	if err := model.QueryGetSong(db, params["song"], &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}
	
	for _, bk := range selectedSongs {
		json.NewEncoder(w).Encode(bk)
  	}
}

func GetGenre(w http.ResponseWriter, req *http.Request) {
	db := model.InitDB(dbpath)
	defer db.Close()

	params := mux.Vars(req)
	selectedSongs := make([]model.Songs, 0)

	if err := model.QueryGetGenre(db, params["genre"], &selectedSongs); err != nil {
		log.Fatalf("Error selecting song from the DB (%s)", err)
	}

	for _, bk := range selectedSongs {
		json.NewEncoder(w).Encode(bk)
  	}
}

func main() {
	router := mux.NewRouter()	

	router.HandleFunc("/artist/{artist}", GetArtist).Methods("GET")
	router.HandleFunc("/song/{song}", GetSong).Methods("GET")
	router.HandleFunc("/genre/{genre}", GetGenre).Methods("GET")

	log.Fatal(http.ListenAndServe(":12345", router))
}

