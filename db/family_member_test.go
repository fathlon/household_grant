package db

import (
	"testing"
	"time"

	"github.com/fathlon/household_grant/model"
	"github.com/stretchr/testify/require"
)

func TestCreateFamilyMember(t *testing.T) {
	validDOB := time.Now().AddDate(-20, 0, 0)

	testCases := []struct {
		msg               string
		givenDatastore    *Datastore
		givenFamilyMember model.FamilyMember
		expectedError     error
	}{
		{
			msg: "FAILURE_CASE",
			givenDatastore: &Datastore{
				Members: map[int]model.FamilyMember{
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
			msg:            "SUCCESS_CASE",
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

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// Given:
			oldIndex := MemIndex
			MemIndex = 0
			defer func() { MemIndex = oldIndex }()

			// When:
			result, err := tc.givenDatastore.CreateFamilyMember(tc.givenFamilyMember)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.NotEmpty(t, result.ID)
			}
		})
	}
}
