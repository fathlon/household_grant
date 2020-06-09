package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fathlon/household_grant/db"
)

var (
	datastore *db.Datastore
)

// init will initialize a new datastore
func init() {
	datastore = db.NewDatastore()
}

// DBServer will return the pointer to db datastore
func DBServer() *db.Datastore {
	return datastore
}

// StartServer will define http routes and start server on given port
func StartServer(port int) {
	r := Routes()

	log.Printf("Starting server at port %v\n", port)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), r))
}
