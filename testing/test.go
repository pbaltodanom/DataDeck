package api_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
)

var (
    server   *httptest.Server
    reader   io.Reader //Ignore this for now
    usersUrl string
)

func init() {
    server = httptest.NewServer(Handlers()) //Creating new server with the user handlers

    usersUrl = fmt.Sprintf("%s/artist/Beatles", server.URL) //Grab the address for the API endpoint
}