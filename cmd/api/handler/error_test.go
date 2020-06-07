package handler

import (
	"errors"
	"net/http"
	"strconv"
	"testing"

	"github.com/fathlon/household_grant/service"
	"github.com/stretchr/testify/require"
)

func TestCheckError(t *testing.T) {
	testCases := []struct {
		given        error
		expectedMsg  string
		expectedCode int
	}{
		{
			service.NewValidationError(errors.New("some validation error")),
			"some validation error",
			http.StatusBadRequest,
		},
		{
			errors.New("an unexpected error"),
			"an unexpected error",
			http.StatusInternalServerError,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// When:
			msg, code := CheckError(tc.given)

			// Then:
			require.Equal(t, tc.expectedMsg, msg)
			require.Equal(t, tc.expectedCode, code)
		})
	}
}
