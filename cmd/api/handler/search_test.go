package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/fathlon/household_grant/db"
	"github.com/fathlon/household_grant/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/require"
)

var (
	now  = time.Now()
	eld1 = model.FamilyMember{
		Name:           "Elder1",
		Gender:         "M",
		MaritalStatus:  "Single",
		OccupationType: "Unemployed",
		DOB:            now.AddDate(-67, 0, 0),
	}
	sm1 = model.FamilyMember{
		Name:           "Single Mum",
		Gender:         "F",
		MaritalStatus:  "Single",
		OccupationType: "Employed",
		AnnualIncome:   250000,
		DOB:            now.AddDate(-34, 0, 0),
	}
	ch1 = model.FamilyMember{
		Name:   "Newborn",
		Gender: "M",
		DOB:    now.AddDate(0, -3, 0),
	}

	// single elder no income
	h1 = model.Household{
		ID:      1,
		Type:    "HDB",
		Members: []model.FamilyMember{eld1},
	}
	// single mother with newborn, total income over 250000
	h2 = model.Household{
		ID:      2,
		Type:    "Landed",
		Members: []model.FamilyMember{sm1, ch1},
	}
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		msg         string
		reqParamStr string
		expected    []model.Household
	}{
		{
			msg:         "empty_params",
			reqParamStr: "",
			expected:    []model.Household{},
		},
		{
			msg:         "household_with_members",
			reqParamStr: "?household_size_gt=1",
			expected:    []model.Household{h2},
		},
		{
			msg:         "household_only_without_members",
			reqParamStr: "?household_size_gt=1&whole_household=true",
			expected: []model.Household{
				{
					ID:   2,
					Type: "Landed",
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// Given:
			ds := &db.Datastore{
				Households: map[int]model.Household{
					1: h1,
					2: h2,
				},
			}

			oldDs := datastore
			datastore = ds
			defer func() { datastore = oldDs }()

			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/search"+tc.reqParamStr, nil)
			require.NoError(t, err)

			router := mux.NewRouter()

			// When:
			router.HandleFunc("/search", Search)
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

func TestSearch_Error(t *testing.T) {
	testCases := []struct {
		msg          string
		reqParamStr  string
		expectedCode int
		expectedMsg  string
	}{
		{
			msg:          "conflicting_params",
			reqParamStr:  "?annual_income_gt=10000&annual_income=15000",
			expectedCode: http.StatusBadRequest,
			expectedMsg:  "conflicting comparison for search",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// Given:
			w := httptest.NewRecorder()
			r, err := http.NewRequest(http.MethodGet, "/search"+tc.reqParamStr, nil)
			require.NoError(t, err)

			router := mux.NewRouter()

			// When:
			router.HandleFunc("/search", Search)
			router.ServeHTTP(w, r)

			// Then:
			require.Equal(t, tc.expectedCode, w.Code)
			require.Contains(t, w.Body.String(), tc.expectedMsg)
		})
	}
}

func TestMapSearchOperation(t *testing.T) {
	truePointer := true
	testCases := []struct {
		msg            string
		givenParams    map[string][]string
		expectedError  error
		expectedResult model.SearchOperation
	}{
		{
			msg: "ignore_duplicate_values",
			givenParams: map[string][]string{
				"annual_income_gt": {"51000"},
				"household_size":   {"3", "6"},
				"has_couple":       {"true"},
			},
			expectedError: nil,
			expectedResult: model.SearchOperation{
				AnnualIncomeGT: 51000,
				HouseholdSize:  3,
				HasCouple:      &truePointer,
			},
		},
		{
			msg: "ignore_unknown_fields",
			givenParams: map[string][]string{
				"age_lt":         {"50"},
				"cant_find_me":   {"hi"},
				"household_size": {"6"},
			},
			expectedError: nil,
			expectedResult: model.SearchOperation{
				AgeLT:         50,
				HouseholdSize: 6,
			},
		},
		{
			msg: "wrong_field_value_int",
			givenParams: map[string][]string{
				"age_lt": {"wrong_type"},
			},
			expectedError:  errors.New("invalid type of value"),
			expectedResult: model.SearchOperation{},
		},
		{
			msg: "wrong_field_value_bool",
			givenParams: map[string][]string{
				"has_couple": {"wrong_type"},
			},
			expectedError:  errors.New("invalid type of value"),
			expectedResult: model.SearchOperation{},
		},
		{
			msg: "skip_key_with_no_value",
			givenParams: map[string][]string{
				"household_income":    {"230000"},
				"household_size":      {},
				"has_children_by_age": {"16"},
			},
			expectedError: nil,
			expectedResult: model.SearchOperation{
				HouseholdIncome:  230000,
				HasChildrenByAge: 16,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			result, err := mapSearchOperation(tc.givenParams)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.EqualValues(t, tc.expectedResult, result)
			}
		})
	}
}
