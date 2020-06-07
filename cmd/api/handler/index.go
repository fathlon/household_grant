package handler

import (
	"fmt"
	"net/http"
)

// Index is the default index page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index")
}
