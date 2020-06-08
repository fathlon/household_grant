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
				households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
				},
			},
			givenHousehold: model.Household{Type: "CONDO"},
			expectedError:  ErrHouseholdDuplicateID,
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

func TestRetrieveHousehold(t *testing.T) {
	testCases := []struct {
		givenDatastore *Datastore
		givenID        int
		expectedError  error
	}{
		{
			givenDatastore: &Datastore{
				households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
					2: {ID: 2, Type: "Landed"},
				},
			},
			givenID:       2,
			expectedError: nil,
		},
		{
			givenDatastore: NewDatastore(),
			givenID:        1,
			expectedError:  ErrHouseholdNotFound,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// When:
			result, err := tc.givenDatastore.RetrieveHousehold(tc.givenID)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.Equal(t, tc.givenID, result.ID)
			}
		})
	}
}
