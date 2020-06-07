package handler

import "net/http"

// Routes defines all accessible http requests
func Routes() {
	http.HandleFunc("/", Index)

	http.HandleFunc("/household/create", CreateHousehold)
}
