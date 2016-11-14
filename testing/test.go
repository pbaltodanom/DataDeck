package api_test

import (
    "DataDeck"
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

var (
    server   *httptest.Server
    reader   io.Reader //Ignore this for now
    usersUrl string
)

func init() {
    server = httptest.NewServer(api.Handlers()) //Creating new server with the user handlers

    usersUrl = fmt.Sprintf("%s/artist/Beatles", server.URL) //Grab the address for the API endpoint
}

