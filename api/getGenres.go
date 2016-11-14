package api

//---------------- IMPORTS ----------------
import (
	"log"
	"database/sql"
)

//Extra #2: Make a function in the API return a list of the genres, and the number of songs and the total length of all the songs by genre
func QueryGetGenres(db *sql.DB, SongsArray *[]Songs) error {
	rows, err := db.Query("SELECT distinct genres.name, COUNT(songs.ID), SUM(songs.length) FROM songs INNER JOIN genres  ON songs.genre = genres.ID GROUP BY  songs.genre")
	if err != nil {
		log.Fatal(err)
	}
	
	//---------------- VARIABLES WHERE THE DATA RETURNED BY THE QUERY ARE KEPT  ----------------
	var resultSongGenre string
	var resultNumberSongs, resultTotalLength uint
	
	//nil slice (type SONGS)
	selectedSongs := Songs{}
	for rows.Next() {
        if err := rows.Scan(&resultSongGenre, &resultNumberSongs, &resultTotalLength); err != nil {
			return err
		}
		//Assign values returned by Scan to slice of type SONGS
		selectedSongs.GenreName = resultSongGenre
		selectedSongs.NumberSongs = resultNumberSongs
		selectedSongs.TotalLength = resultTotalLength	
		
		//Append the slice in the SongsArray
		*SongsArray = append(*SongsArray, selectedSongs)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
	
	return nil
}