package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Routes defines all accessible http requests
func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", Index)

	hr := r.PathPrefix("/households").Subrouter()
	hr.HandleFunc("", CreateHousehold).Methods(http.MethodPost)
	hr.HandleFunc("/{id}/familymember", AddFamilyMember).Methods(http.MethodPost)

	return r
}
