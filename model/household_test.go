package model

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateHousehold(t *testing.T) {
	testCases := []struct {
		given    Household
		expected error
	}{
		{Household{Type: "HDB"}, nil},
		{Household{Type: "Landed"}, nil},
		{Household{Type: "Condominium"}, nil},
		{Household{Type: "Special"}, ErrHouseholdTypeInvalid},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			// When:
			result := tc.given.ValidateHousehold()

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}
