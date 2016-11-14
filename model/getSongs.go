package model

//---------------- IMPORTS ---------------- 
import (
	"log"
	"database/sql"
)

//Extra #1: Make a function in the API return songs by length, which we would like to search by passing a minimum and maximum length.
func QueryGetSongs(db *sql.DB, min string, max string, SongsArray *[]Songs) error {
	rows, err := db.Query("SELECT song, length FROM songs WHERE length BETWEEN ? AND ? ORDER BY length;", min, max)
	if err != nil {
		log.Fatal(err)
	}
	
	//---------------- VARIABLES WHERE THE DATA RETURNED BY THE QUERY ARE KEPT  ----------------
	var resultSongName string
	var resultSongLength uint
	
	//nil slice (type SONGS)
	selectedSongs := Songs{}
	for rows.Next() {
        if err := rows.Scan(&resultSongName, &resultSongLength); err != nil {
			return err
		}
		//Assign values returned by Scan to slice of type SONGS
		selectedSongs.Name = resultSongName
		selectedSongs.Length = resultSongLength
		
		//Append the slice in the SongsArray
		*SongsArray = append(*SongsArray, selectedSongs)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
	
	return nil
}