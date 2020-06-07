package db

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddHousehold(t *testing.T) {
	testCases := []struct {
		givenDatastore *Datastore
		givenHousehold Household
		expectedError  error
	}{
		{
			givenDatastore: &Datastore{
				store: map[int]Household{
					1: {ID: 1, Type: "HDB"},
				},
			},
			givenHousehold: Household{Type: "CONDO"},
			expectedError:  ErrDuplicateID,
		},
		{
			givenDatastore: NewDatastore(),
			givenHousehold: Household{Type: "CONDO"},
			expectedError:  nil,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// When:
			result, err := tc.givenDatastore.AddHousehold(tc.givenHousehold)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.NotEmpty(t, result.ID)
			}
		})
	}
}
