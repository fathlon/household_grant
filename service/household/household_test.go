package household

import (
	"testing"
	"time"

	"github.com/fathlon/household_grant/db"
	"github.com/fathlon/household_grant/model"
	"github.com/fathlon/household_grant/service"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	testCases := []struct {
		msg            string
		givenDatastore *db.Datastore
		givenHousehold model.Household
		expectedError  error
	}{
		{
			msg:            "FAILURE_CASE",
			givenDatastore: db.NewDatastore(),
			givenHousehold: model.Household{},
			expectedError:  service.NewValidationError(model.ErrHouseholdTypeInvalid),
		},
		{
			msg:            "SUCCESS_CASE",
			givenDatastore: db.NewDatastore(),
			givenHousehold: model.Household{Type: "Landed"},
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()

			// When:
			result, err := Create(tc.givenDatastore, tc.givenHousehold)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.NotZero(t, result.ID)
				require.Equal(t, tc.givenHousehold.Type, result.Type)
			}
		})
	}
}
func TestAddMember(t *testing.T) {
	testCases := []struct {
		msg              string
		givenDatastore   *db.Datastore
		givenHouseholdID int
		givenMember      model.FamilyMember
	}{
		{
			msg: "MEMBER_WITH_ID",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
				},
				Members: map[int]model.FamilyMember{
					1: {
						ID:             1,
						Name:           "Jackie",
						Gender:         "M",
						OccupationType: "Unemployed",
						MaritalStatus:  "Married",
						DOB:            time.Now(),
					},
				},
			},
			givenHouseholdID: 1,
			givenMember: model.FamilyMember{
				ID:             2,
				Name:           "Alexia",
				Gender:         "F",
				OccupationType: "Student",
				MaritalStatus:  "Single",
				DOB:            time.Now(),
			},
		},
		{
			msg: "MEMBER_WITHOUT_ID",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
				},
				Members: map[int]model.FamilyMember{},
			},
			givenHouseholdID: 1,
			givenMember: model.FamilyMember{
				Name:           "Alexia",
				Gender:         "F",
				OccupationType: "Student",
				MaritalStatus:  "Single",
				DOB:            time.Now(),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result, err := AddMember(tc.givenDatastore, tc.givenHouseholdID, tc.givenMember)

			// Then:
			require.NoError(t, err)
			require.NotNil(t, result.ID)
			// assert member is added to DB
			require.NotNil(t, tc.givenDatastore.Members[result.ID])

			// assert member is added to household
			actual := tc.givenDatastore.Households[tc.givenHouseholdID]
			require.True(t, actual.MemberExists(result.ID))
		})
	}
}

func TestAddMember_Error(t *testing.T) {
	testCases := []struct {
		msg              string
		givenDatastore   *db.Datastore
		givenHouseholdID int
		givenMember      model.FamilyMember
		expectedError    *service.ValidationError
	}{
		{
			msg:              "INVALID_MEMBER",
			givenDatastore:   db.NewDatastore(),
			givenHouseholdID: 1,
			givenMember:      model.FamilyMember{},
			expectedError:    service.NewValidationError(model.ErrFamilyMemberNameInvalid),
		},
		{
			msg:              "INVALID_HOUSEHOLD",
			givenDatastore:   db.NewDatastore(),
			givenHouseholdID: 1,
			givenMember: model.FamilyMember{
				Name:           "Jackie",
				Gender:         "M",
				OccupationType: "Unemployed",
				MaritalStatus:  "Married",
				DOB:            time.Now(),
			},
			expectedError: service.NewValidationError(db.ErrHouseholdNotFound),
		},
		{
			msg:              "INVALID_HOUSEHOLD",
			givenDatastore:   db.NewDatastore(),
			givenHouseholdID: 1,
			givenMember: model.FamilyMember{
				Name:           "Jackie",
				Gender:         "M",
				OccupationType: "Unemployed",
				MaritalStatus:  "Married",
				DOB:            time.Now(),
			},
			expectedError: service.NewValidationError(db.ErrHouseholdNotFound),
		},
		{
			msg: "DUPLICATE_MEMBER_ID",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
				},
				Members: map[int]model.FamilyMember{
					1: {
						ID:             1,
						Name:           "Jackie",
						Gender:         "M",
						OccupationType: "Unemployed",
						MaritalStatus:  "Married",
						DOB:            time.Now(),
					},
				},
			},
			givenHouseholdID: 1,
			givenMember: model.FamilyMember{
				Name:           "Jackie",
				Gender:         "M",
				OccupationType: "Unemployed",
				MaritalStatus:  "Married",
				DOB:            time.Now(),
			},
			expectedError: service.NewValidationError(db.ErrFamilyMemberDuplicateID),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// Given:
			oldIndex := db.MemIndex
			db.MemIndex = 0
			defer func() { db.MemIndex = oldIndex }()

			// When:
			_, err := AddMember(tc.givenDatastore, tc.givenHouseholdID, tc.givenMember)

			// Then:
			require.Equal(t, tc.expectedError, err)
		})
	}
}
