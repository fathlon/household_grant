package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fathlon/household_grant/db"
)

var datastore *db.Datastore

// Init will initialize a new datastore
func Init() {
	datastore = db.NewDatastore()
}

func main() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Index is the default index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}
