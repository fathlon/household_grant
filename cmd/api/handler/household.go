package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/fathlon/household_grant/model"
	"github.com/fathlon/household_grant/service/household"
)

// CreateHousehold is the handler for creating new household
func CreateHousehold(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var newHousehold model.Household
	if err := json.Unmarshal(reqBody, &newHousehold); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	result, err := household.Create(DBServer(), newHousehold)
	if err != nil {
		errMsg, errCode := CheckError(err)
		http.Error(w, errMsg, errCode)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
