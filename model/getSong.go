package model

import (
	"log"
	"database/sql"
)

func QueryGetSong(db *sql.DB, song string, SongsArray *[]Songs) error {
	rows, err := db.Query("SELECT * FROM songs WHERE song = ?", song)
	if err != nil {
		log.Fatal(err)
	}

	selectedSongs := Songs{}
	for rows.Next() {
        err := rows.Scan(&selectedSongs.ID, &selectedSongs.Artist, &selectedSongs.Name, &selectedSongs.GenreName, &selectedSongs.Length)
        if err != nil {
  			log.Fatal(err)
		}
		*SongsArray = append(*SongsArray, selectedSongs)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
	
	return nil
}