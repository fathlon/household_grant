package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/fathlon/household_grant/db"
	"github.com/fathlon/household_grant/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

func TestCreateHousehold(t *testing.T) {
	// Given:
	reqHousehold := model.Household{Type: "Landed"}
	reqData, err := json.Marshal(reqHousehold)
	require.NoError(t, err)

	w := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodPost, "/households", bytes.NewBuffer(reqData))
	require.NoError(t, err)

	router := mux.NewRouter()

	// When:
	router.HandleFunc("/households", CreateHousehold)
	router.ServeHTTP(w, r)

	// Then:
	require.Equal(t, http.StatusCreated, w.Code)
	require.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var result model.Household
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &result))
	require.NotEmpty(t, result.ID)
	require.Equal(t, reqHousehold.Type, result.Type)

}

func TestCreateHousehold_Error(t *testing.T) {
	testCases := []struct {
		msg          string
		givenReqStr  string
		expectedCode int
		expectedMsg  string
	}{
		{
			msg:          "invalid_json",
			givenReqStr:  "{\"type\": \"invalid json\"",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "error parsing json:",
		},
		{
			msg:          "invalid_field_value",
			givenReqStr:  "{\"type\": \"special house\"}",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  model.ErrHouseholdTypeInvalid.Error(),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// Given:
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodPost, "/households", bytes.NewBufferString(tc.givenReqStr))
			require.NoError(t, err)

			router := mux.NewRouter()

			// When:
			router.HandleFunc("/households", CreateHousehold)
			router.ServeHTTP(w, r)

			// Then:
			require.Equal(t, tc.expectedCode, w.Code)
			require.Contains(t, w.Body.String(), tc.expectedMsg)
		})
	}
}

func TestAddFamilyMember(t *testing.T) {
	mb1 := model.FamilyMember{
		ID:             1,
		Name:           "Alexia",
		Gender:         "M",
		MaritalStatus:  "Single",
		OccupationType: "Student",
		DOB:            time.Now(),
	}
	testCases := []struct {
		msg               string
		givenDatastore    *db.Datastore
		givenHouseholdID  string
		givenFamilyMember model.FamilyMember
	}{
		{
			msg: "new_member",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {
						ID:   1,
						Type: "landed",
					},
				},
				Members: make(map[int]model.FamilyMember),
			},
			givenHouseholdID: "1",
			givenFamilyMember: model.FamilyMember{
				Name:           "Alexia",
				Gender:         "M",
				MaritalStatus:  "Single",
				OccupationType: "Student",
				DOB:            time.Now(),
			},
		},
		{
			msg: "existing_member",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {
						ID:      1,
						Type:    "landed",
						Members: []model.FamilyMember{mb1},
					},
				},
				Members: make(map[int]model.FamilyMember),
			},
			givenHouseholdID:  "1",
			givenFamilyMember: mb1,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// Given:
			oldDs := datastore
			datastore = tc.givenDatastore
			defer func() { datastore = oldDs }()

			reqData, err := json.Marshal(tc.givenFamilyMember)
			require.NoError(t, err)

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/households/%v/familymember", tc.givenHouseholdID), bytes.NewBuffer(reqData))
			require.NoError(t, err)

			router := mux.NewRouter()

			// When:
			router.HandleFunc("/households/{id}/familymember", AddFamilyMember)
			router.ServeHTTP(w, r)

			// Then:
			require.Equal(t, http.StatusOK, w.Code)
			require.Equal(t, "application/json", w.Header().Get("Content-Type"))

			var result model.FamilyMember
			require.NoError(t, json.Unmarshal(w.Body.Bytes(), &result))
			require.NotEmpty(t, result.ID)
		})
	}
}

func TestAddFamilyMember_Error(t *testing.T) {
	testCases := []struct {
		msg              string
		givenHouseholdID string
		givenReqStr      string
		expectedCode     int
		expectedMsg      string
	}{
		{
			msg:              "invalid_path_variable",
			givenHouseholdID: "one",
			givenReqStr:      "",
			expectedCode:     http.StatusBadRequest,
			expectedMsg:      "invalid path variable",
		},
		{
			msg:              "invalid_json",
			givenHouseholdID: "1",
			givenReqStr:      "{ invalid: json }",
			expectedCode:     http.StatusBadRequest,
			expectedMsg:      "error parsing json:",
		},
		{
			msg:              "-",
			givenHouseholdID: "1",
			givenReqStr: `{
				"name": "Alexia",
				"gender": "Trans",
				"marital_status": "Single",
				"spouse": "",
				"occupation_type": "Student",
				"annual_income": 0,
				"dob": "1990-06-09T09:05:18+08:00"
			}`,
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "invalid gender",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// Given:
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodPost, fmt.Sprintf("/households/%v/familymember", tc.givenHouseholdID), bytes.NewBufferString(tc.givenReqStr))
			require.NoError(t, err)

			router := mux.NewRouter()

			// When:
			router.HandleFunc("/households/{id}/familymember", AddFamilyMember)
			router.ServeHTTP(w, r)

			// Then:
			require.Equal(t, tc.expectedCode, w.Code)
			require.Contains(t, w.Body.String(), tc.expectedMsg)
		})
	}
}

func TestRetrieveHouseholds(t *testing.T) {
	h1 := model.Household{
		ID:   1,
		Type: "Landed",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Jack"},
			{ID: 2, Name: "Beanstalk"},
		},
	}
	h2 := model.Household{
		ID:   2,
		Type: "HDB",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Cinderella"},
		},
	}

	testCases := []struct {
		msg            string
		givenDatastore *db.Datastore
		expected       []model.Household
	}{
		{
			msg:            "empty",
			givenDatastore: db.NewDatastore(),
			expected:       []model.Household{},
		},
		{
			msg: "not_empty",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: h1,
					2: h2,
				},
			},
			expected: []model.Household{h1, h2},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// Given:
			oldDs := datastore
			datastore = tc.givenDatastore
			defer func() { datastore = oldDs }()

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/households", nil)
			require.NoError(t, err)

			router := mux.NewRouter()

			// When:
			router.HandleFunc("/households", RetrieveHouseholds)
			router.ServeHTTP(w, r)

			// Then:
			require.Equal(t, http.StatusOK, w.Code)
			require.Equal(t, "application/json", w.Header().Get("Content-Type"))

			var result []model.Household
			require.NoError(t, json.Unmarshal(w.Body.Bytes(), &result))
			require.Equal(t, len(tc.expected), len(result))
			require.ElementsMatch(t, tc.expected, result)
		})
	}
}
