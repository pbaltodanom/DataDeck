package model

import (
	"log"
	"database/sql"
)

func QueryGetSong(db *sql.DB, song string, SongsArray *[]Songs) error {
	rows, err := db.Query("SELECT song, artist, genres.name, length FROM songs INNER JOIN genres ON songs.genre = genres.ID WHERE song = ?", song)
	if err != nil {
		log.Fatal(err)
	}

	var resultSongArtist, resultSongName, resultSongGenre string
	var resultSongLength uint

	selectedSongs := Songs{}
	for rows.Next() {
        if err := rows.Scan(&resultSongName, &resultSongArtist, &resultSongGenre, &resultSongLength); err != nil {
			return err
		}

		selectedSongs.Artist = resultSongArtist
		selectedSongs.Name = resultSongName
		selectedSongs.GenreName = resultSongGenre
		selectedSongs.Length = resultSongLength
		*SongsArray = append(*SongsArray, selectedSongs)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
	
	return nil
}