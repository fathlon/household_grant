package db

import (
	"strconv"
	"testing"

	"github.com/fathlon/household_grant/model"
	"github.com/stretchr/testify/require"
)

func TestAddHousehold(t *testing.T) {
	testCases := []struct {
		givenDatastore *Datastore
		givenHousehold model.Household
		expectedError  error
	}{
		{
			givenDatastore: &Datastore{
				store: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
				},
			},
			givenHousehold: model.Household{Type: "CONDO"},
			expectedError:  ErrDuplicateID,
		},
		{
			givenDatastore: NewDatastore(),
			givenHousehold: model.Household{Type: "CONDO"},
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
