package main

//TEST - DEVELOPED IN GOLANG 1.6.3

//---------------- IMPORTS ---------------- 
import (
	"DataDeck/api"
	//"github.com/pbaltodanom/DataDeck/api"
	"log"	
	"net/http"		

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//ListenAndServe starts an HTTP server with a given address and handler.	
	log.Fatal(http.ListenAndServe(":12345", api.Handlers()))
}

