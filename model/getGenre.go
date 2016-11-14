package model

import (
	"log"
	"database/sql"
)

func QueryGetGenre(db *sql.DB, genre string, SongsArray *[]Songs) error {
	rows, err := db.Query("SELECT songs.id, artist, song, genres.name, length FROM songs INNER JOIN genres ON songs.genre=genres.ID WHERE genres.Name = ?", genre)
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