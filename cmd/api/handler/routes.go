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
	hr.HandleFunc("", RetrieveHouseholds).Methods(http.MethodGet)
	hr.HandleFunc("/{id}", RetrieveHousehold).Methods(http.MethodGet)
	hr.HandleFunc("/{id}", DeleteHousehold).Methods(http.MethodDelete)
	hr.HandleFunc("/{id}/familymember", AddFamilyMember).Methods(http.MethodPost)
	hr.HandleFunc("/{id}/familymember/{fmid}", DeleteFamilyMember).Methods(http.MethodDelete)

	r.HandleFunc("/search", Search).Methods(http.MethodGet)

	return r
}
