package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fathlon/household_grant/cmd/api/handler"
	"github.com/fathlon/household_grant/db"
)

var (
	datastore *db.Datastore

	port = 8080
)

// Init will initialize a new datastore
func Init() {
	datastore = db.NewDatastore()
}

func main() {
	http.HandleFunc("/", handler.Index)

	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}
