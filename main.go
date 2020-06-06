package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Index)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Index is the default index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}
