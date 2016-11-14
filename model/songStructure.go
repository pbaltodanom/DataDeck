package model

// Person is the database model for a person
type Songs struct {
	//ID  uint `json:"ID,omitempty"` never returns in the row
	Artist string `json:"Artist,omitempty"`
	Name  string `json:"Name,omitempty"`
	GenreName  string `json:"Genre,omitempty"`
	Length  uint `json:"Length,omitempty"`
}