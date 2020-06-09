package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fathlon/household_grant/model"
	"github.com/fathlon/household_grant/service/household"
	"github.com/gorilla/mux"
)

// CreateHousehold is the handler for creating new household
func CreateHousehold(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var newHousehold model.Household
	if err := json.Unmarshal(reqBody, &newHousehold); err != nil {
		http.Error(w, fmt.Sprintf("error parsing json: %v", err.Error()), http.StatusBadRequest)
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
		http.Error(w, fmt.Sprintf("error parsing json: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
}

// AddFamilyMember is the handler for adding family member to a household
func AddFamilyMember(w http.ResponseWriter, r *http.Request) {
	pathID := mux.Vars(r)["id"]
	householdID, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path variable: %v", pathID), http.StatusBadRequest)
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var member model.FamilyMember
	if err := json.Unmarshal(reqBody, &member); err != nil {
		http.Error(w, fmt.Sprintf("error parsing json: %v", err.Error()), http.StatusBadRequest)
		return
	}

	result, err := household.AddMember(DBServer(), householdID, member)
	if err != nil {
		errMsg, errCode := CheckError(err)
		http.Error(w, errMsg, errCode)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing json: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// RetrieveHouseholds returns all households in the db
func RetrieveHouseholds(w http.ResponseWriter, r *http.Request) {
	result := household.RetrieveAll(DBServer())

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing json: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// RetrieveHousehold returns the household of the given id
func RetrieveHousehold(w http.ResponseWriter, r *http.Request) {
	pathID := mux.Vars(r)["id"]
	householdID, err := strconv.Atoi(pathID)
	if err != nil {
		http.Error(w, fmt.Sprintf("invalid path variable: %v", pathID), http.StatusBadRequest)
		return
	}

	result, err := household.Retrieve(DBServer(), householdID)
	if err != nil {
		errMsg, errCode := CheckError(err)
		http.Error(w, errMsg, errCode)
		return
	}

	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("error parsing json: %v", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
