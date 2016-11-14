package model

//---------------- IMPORTS ---------------- 
import (
	"log"
	"database/sql"
)

//Access the genre
//Retrieve the data that meets the query specifications
func QueryGetGenre(db *sql.DB, genre string, SongsArray *[]Songs) error {
	//Query executes a query that returns rows
	rows, err := db.Query("SELECT song, artist, genres.name, length FROM songs INNER JOIN genres ON songs.genre=genres.ID WHERE genres.Name = ?", genre)
	if err != nil {
		log.Fatal(err)
	}
	
	//---------------- VARIABLES WHERE THE DATA RETURNED BY THE QUERY ARE KEPT  ----------------
	var resultSongArtist, resultSongName, resultSongGenre string
	var resultSongLength uint
	
	//nil slice (type SONGS)
	selectedSongs := Songs{}
	for rows.Next() {
		//For each row
        if err := rows.Scan(&resultSongName, &resultSongArtist, &resultSongGenre, &resultSongLength); err != nil {
			return err
		}
		//Assign values returned by Scan to slice of type SONGS
		selectedSongs.Artist = resultSongArtist
		selectedSongs.Name = resultSongName
		selectedSongs.GenreName = resultSongGenre
		selectedSongs.Length = resultSongLength		

		//Append the slice in the SongsArray
		*SongsArray = append(*SongsArray, selectedSongs)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
	
	return nil
}