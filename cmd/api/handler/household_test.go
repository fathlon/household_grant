package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/fathlon/household_grant/model"
	"github.com/stretchr/testify/require"
)

func TestCreateHousehold(t *testing.T) {
	// Given:
	reqHousehold := model.Household{Type: "Landed"}
	reqData, err := json.Marshal(reqHousehold)
	require.NoError(t, err)

	r := httptest.NewRequest(http.MethodPost, "/household/create", bytes.NewBuffer(reqData))
	w := httptest.NewRecorder()

	// When:
	CreateHousehold(w, r)

	// Then:
	require.Equal(t, "application/json", w.Header().Get("Content-Type"))
	require.Equal(t, http.StatusOK, w.Code)

	var result model.Household
	require.NoError(t, json.Unmarshal(w.Body.Bytes(), &result))
	require.NotEmpty(t, result.ID)
	require.Equal(t, reqHousehold.Type, result.Type)

}

func TestCreateHousehold_Error(t *testing.T) {
	testCases := []struct {
		givenHHTTPMethod string
		givenReqStr      string
		expectedCode     int
		expectedMsg      string
	}{
		{
			givenHHTTPMethod: http.MethodGet,
			givenReqStr:      "",
			expectedCode:     http.StatusMethodNotAllowed,
			expectedMsg:      http.StatusText(http.StatusMethodNotAllowed) + "\n",
		},
		{
			givenHHTTPMethod: http.MethodPost,
			givenReqStr:      "{\"type\": \"invalid json\"",
			expectedCode:     http.StatusBadRequest,
			expectedMsg:      http.StatusText(http.StatusBadRequest) + "\n",
		},
		{
			givenHHTTPMethod: http.MethodPost,
			givenReqStr:      "{\"type\": \"special house\"}",
			expectedCode:     http.StatusBadRequest,
			expectedMsg:      model.ErrHouseholdTypeInvalid.Error() + "\n",
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			// Given:
			r := httptest.NewRequest(tc.givenHHTTPMethod, "/household/create", bytes.NewBufferString(tc.givenReqStr))
			w := httptest.NewRecorder()

			// When:
			CreateHousehold(w, r)

			// Then:
			require.Equal(t, tc.expectedCode, w.Code)
			require.Equal(t, tc.expectedMsg, w.Body.String())
		})
	}
}
