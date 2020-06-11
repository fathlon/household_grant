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
			msg:            "failure",
			givenDatastore: db.NewDatastore(),
			givenHousehold: model.Household{},
			expectedError:  service.NewValidationError(model.ErrHouseholdTypeInvalid),
		},
		{
			msg:            "success",
			givenDatastore: db.NewDatastore(),
			givenHousehold: model.Household{
				Type: "Landed",
				Members: []model.FamilyMember{
					{ID: 1, Name: "Sleeping Beauty"},
					{ID: 2, Name: "Snow White"},
				},
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			result, err := Create(tc.givenDatastore, tc.givenHousehold)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.NotZero(t, result.ID)
				require.Equal(t, tc.givenHousehold.Type, result.Type)
				require.Empty(t, result.Members)
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
			msg: "member_with_id",
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
			msg: "member_without_id",
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
			msg:              "invalid_member",
			givenDatastore:   db.NewDatastore(),
			givenHouseholdID: 1,
			givenMember:      model.FamilyMember{},
			expectedError:    service.NewValidationError(model.ErrFamilyMemberNameInvalid),
		},
		{
			msg:              "invalid_household",
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
			msg: "duplicate_member_id",
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
				Name:           "Harry",
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

func TestRetrieveAll(t *testing.T) {
	h1 := model.Household{
		ID:   1,
		Type: "Landed",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Jack"},
			{ID: 2, Name: "Beanstalk"},
		},
	}
	h2 := model.Household{
		ID:   2,
		Type: "HDB",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Cinderella"},
		},
	}

	testCases := []struct {
		msg            string
		givenDatastore *db.Datastore
		expected       []model.Household
	}{
		{
			msg:            "empty",
			givenDatastore: db.NewDatastore(),
			expected:       []model.Household{},
		},
		{
			msg: "not_empty",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: h1,
					2: h2,
				},
			},
			expected: []model.Household{h1, h2},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			result := RetrieveAll(tc.givenDatastore)

			// Then:
			require.Equal(t, len(tc.expected), len(result))
			require.ElementsMatch(t, tc.expected, result)
		})
	}
}

func TestRetrieve(t *testing.T) {
	h1 := model.Household{
		ID:   1,
		Type: "Landed",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Jack"},
			{ID: 2, Name: "Beanstalk"},
		},
	}
	h2 := model.Household{
		ID:   2,
		Type: "HDB",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Cinderella"},
		},
	}

	ds := db.Datastore{
		Households: map[int]model.Household{
			1: h1,
			2: h2,
		},
	}

	testCases := []struct {
		msg               string
		givenHouseholdID  int
		expectedError     error
		expectedHousehold model.Household
	}{
		{
			msg:               "not_found",
			givenHouseholdID:  5,
			expectedError:     service.NewValidationError(db.ErrHouseholdNotFound),
			expectedHousehold: model.Household{},
		},
		{
			msg:               "found",
			givenHouseholdID:  2,
			expectedHousehold: h2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			result, err := Retrieve(&ds, tc.givenHouseholdID)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.EqualValues(t, tc.expectedHousehold, result)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	h1 := model.Household{
		ID:   1,
		Type: "Landed",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Jack"},
			{ID: 2, Name: "Beanstalk"},
		},
	}
	h2 := model.Household{
		ID:   2,
		Type: "HDB",
		Members: []model.FamilyMember{
			{ID: 1, Name: "Cinderella"},
		},
	}

	ds := db.Datastore{
		Households: map[int]model.Household{
			1: h1,
			2: h2,
		},
	}

	testCases := []struct {
		msg               string
		givenHouseholdID  int
		expectedError     error
		expectedHousehold model.Household
	}{
		{
			msg:               "not_found",
			givenHouseholdID:  5,
			expectedError:     service.NewValidationError(db.ErrHouseholdNotFound),
			expectedHousehold: model.Household{},
		},
		{
			msg:               "success",
			givenHouseholdID:  2,
			expectedHousehold: h2,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			result, err := Delete(&ds, tc.givenHouseholdID)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.EqualValues(t, tc.expectedHousehold, result)

				_, exists := ds.Households[tc.givenHouseholdID]
				require.False(t, exists)
			}
		})
	}
}

func TestDeleteMember(t *testing.T) {
	mb1 := model.FamilyMember{
		ID:   1,
		Name: "Jackie",
	}

	testCases := []struct {
		msg              string
		givenDatastore   *db.Datastore
		givenHouseholdID int
		givenMemberID    int
		expected         model.FamilyMember
	}{
		{
			msg: "success",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {
						ID:      1,
						Type:    "HDB",
						Members: []model.FamilyMember{mb1},
					},
				},
				Members: map[int]model.FamilyMember{
					1: mb1,
				},
			},
			givenHouseholdID: 1,
			givenMemberID:    1,
			expected:         mb1,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			result, err := DeleteMember(tc.givenDatastore, tc.givenHouseholdID, tc.givenMemberID)

			// Then:
			require.NoError(t, err)
			require.EqualValues(t, tc.expected, result)
		})
	}
}

func TestDeleteMember_Error(t *testing.T) {
	testCases := []struct {
		msg              string
		givenDatastore   *db.Datastore
		givenHouseholdID int
		givenMemberID    int
		expectedError    *service.ValidationError
	}{
		{
			msg:              "household_not_found",
			givenDatastore:   db.NewDatastore(),
			givenHouseholdID: 1,
			givenMemberID:    1,
			expectedError:    service.NewValidationError(db.ErrHouseholdNotFound),
		},
		{
			msg: "member_not_exists_in_household",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
				},
				Members: map[int]model.FamilyMember{
					1: {
						ID:   1,
						Name: "Jackie",
					},
				},
			},
			givenHouseholdID: 1,
			givenMemberID:    1,
			expectedError:    service.NewValidationError(model.ErrHouseholdFamilyMemberNotExists),
		},
		{
			msg: "member_not_exists_master_map",
			givenDatastore: &db.Datastore{
				Households: map[int]model.Household{
					1: {
						ID:   1,
						Type: "HDB",
						Members: []model.FamilyMember{
							{
								ID:   1,
								Name: "Jackie",
							},
						},
					},
				},
				Members: map[int]model.FamilyMember{},
			},
			givenHouseholdID: 1,
			givenMemberID:    1,
			expectedError:    service.NewValidationError(db.ErrFamilyMemberNotFound),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			_, err := DeleteMember(tc.givenDatastore, tc.givenHouseholdID, tc.givenMemberID)

			// Then:
			require.Equal(t, tc.expectedError, err)
		})
	}
}
