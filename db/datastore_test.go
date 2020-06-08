package db

import (
	"strconv"
	"testing"
	"time"

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

func TestAddFamilyMember(t *testing.T) {
	validDOB := time.Now().AddDate(-20, 0, 0)

	testCases := []struct {
		givenDatastore    *Datastore
		givenFamilyMember model.FamilyMember
		expectedError     error
	}{
		{
			givenDatastore: &Datastore{
				members: map[int]model.FamilyMember{
					1: {
						ID:             1,
						Name:           "Jackie",
						Gender:         "M",
						OccupationType: "Unemployed",
						MaritalStatus:  "Married",
						DOB:            validDOB,
					},
				},
			},
			givenFamilyMember: model.FamilyMember{
				Name:           "Jackie",
				Gender:         "M",
				OccupationType: "Unemployed",
				MaritalStatus:  "Married",
				DOB:            validDOB,
			},
			expectedError: ErrFamilyMemberDuplicateID,
		},
		{
			givenDatastore: NewDatastore(),
			givenFamilyMember: model.FamilyMember{
				Name:           "Jackie",
				Gender:         "M",
				OccupationType: "Unemployed",
				MaritalStatus:  "Married",
				DOB:            validDOB,
			},
			expectedError: nil,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// When:
			result, err := tc.givenDatastore.AddFamilyMember(tc.givenFamilyMember)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.NotEmpty(t, result.ID)
			}
		})
	}
}
