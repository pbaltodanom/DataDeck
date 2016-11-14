package model

const (
	SongTableName = "songs"	

	// Names of every column of the model
	
	SongIDCol = "id"
	SongArtistCol = "artist"
	SongNameCol = "song"
	SongGenreCol = "genre"
	SongGenreNameCol = "name"
	SongLengthCol = "length"
)

// Person is the database model for a person
type Songs struct {
	ID  uint `json:"ID,omitempty"`
	Artist string `json:"Artist,omitempty"`
	Name  string `json:"Name,omitempty"`
	Genre  uint `json:"Genre,omitempty"`
	Length  uint `json:"Length,omitempty"`
	GenreName  string `json:"Genre,omitempty"`
}