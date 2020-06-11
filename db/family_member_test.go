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
			msg: "duplicate_id",
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
				Name:           "Harry",
				Gender:         "M",
				OccupationType: "Unemployed",
				MaritalStatus:  "Married",
				DOB:            validDOB,
			},
			expectedError: ErrFamilyMemberDuplicateID,
		},
		{
			msg: "duplicate_name",
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
			expectedError: ErrFamilyMemberDuplicateName,
		},
		{
			msg:            "success",
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

func TestDeleteFamilyMember(t *testing.T) {
	fm1 := model.FamilyMember{
		ID:   1,
		Name: "Jack",
	}
	fm2 := model.FamilyMember{
		ID:   2,
		Name: "Jill",
	}

	testCases := []struct {
		msg            string
		givenDatastore *Datastore
		givenMemberID  int
		expectedError  error
		expected       model.FamilyMember
	}{
		{
			msg:            "not_found",
			givenDatastore: NewDatastore(),
			givenMemberID:  1,
			expectedError:  ErrFamilyMemberNotFound,
			expected:       model.FamilyMember{},
		},
		{
			msg: "success",
			givenDatastore: &Datastore{
				Members: map[int]model.FamilyMember{
					1: fm1,
					2: fm2,
				},
			},
			givenMemberID: fm1.ID,
			expectedError: nil,
			expected:      fm1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result, err := tc.givenDatastore.DeleteFamilyMember(tc.givenMemberID)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.EqualValues(t, tc.expected, result)

				_, exists := tc.givenDatastore.Members[tc.givenMemberID]
				require.False(t, exists)
			}
		})
	}
}
