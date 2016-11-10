package models

import (
	"fmt"
	"database/sql"
)

const (
	// Name of the table for the song model
	SongTableName = "songs"
	// Names of every column of the model
	
	SongArtistCol = "artist"
	SongNameCol = "song"
	SongGenreCol = "genre"
	SongLengthCol = "length"
)

// Person is the database model for a person
type Songs struct {
	Artist string
	Name  string
	Genre  string
	Length  uint
}

func Query(db *sql.DB, artist string, song string, genre uint, length uint, result *Songs) error {
	row := db.QueryRow(
		fmt.Sprintf(
				"SELECT * FROM %s WHERE %s=? AND %s=? AND %s=?",
			SongTableName,
			SongArtistCol,
			SongNameCol,
			SongGenreCol,
			SongLengthCol,
		),
		artist,
		song,
		genre,
		length,
	)
	var resultSongArtist, resultSongName, resultSongGenre string
	var resultSongLength uint
	if err := row.Scan(&resultSongArtist, &resultSongName, &resultSongGenre, resultSongLength); err != nil {
		return err
	}
	result.Artist = resultSongArtist
	result.Name = resultSongName
	result.Genre = resultSongGenre
	result.Length = resultSongLength
	return nil
}