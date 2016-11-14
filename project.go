package main

//TEST - DEVELOPED IN GOLANG 1.6.3

//---------------- IMPORTS ---------------- 
import (
	"github.com/pbaltodanom/DataDeck/api"
	"fmt"	
	"net/http"		

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Server starting")
	//ListenAndServe starts an HTTP server with a given address and handler.
	//log.Fatal(http.ListenAndServe(":12345", api.Handlers()))
	http.ListenAndServe(":3000", api.Handlers())
}

