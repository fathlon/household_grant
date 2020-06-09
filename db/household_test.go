package db

import (
	"testing"

	"github.com/fathlon/household_grant/model"
	"github.com/stretchr/testify/require"
)

func TestCreateHousehold(t *testing.T) {
	testCases := []struct {
		msg            string
		givenDatastore *Datastore
		givenHousehold model.Household
		expectedError  error
	}{
		{
			msg: "failure",
			givenDatastore: &Datastore{
				Households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
				},
			},
			givenHousehold: model.Household{Type: "CONDO"},
			expectedError:  ErrHouseholdDuplicateID,
		},
		{
			msg:            "success",
			givenDatastore: NewDatastore(),
			givenHousehold: model.Household{Type: "CONDO"},
			expectedError:  nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// Given:
			oldIndex := HseIndex
			HseIndex = 0
			defer func() { HseIndex = oldIndex }()

			// When:
			result, err := tc.givenDatastore.CreateHousehold(tc.givenHousehold)

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
		msg            string
		givenDatastore *Datastore
		givenID        int
		expectedError  error
	}{
		{
			msg: "success",
			givenDatastore: &Datastore{
				Households: map[int]model.Household{
					1: {ID: 1, Type: "HDB"},
					2: {ID: 2, Type: "Landed"},
				},
			},
			givenID:       2,
			expectedError: nil,
		},
		{
			msg:            "failure",
			givenDatastore: NewDatastore(),
			givenID:        1,
			expectedError:  ErrHouseholdNotFound,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
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

func TestRetrieveHouseholds(t *testing.T) {
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
		givenDatastore *Datastore
		expected       []model.Household
	}{
		{
			msg:            "empty",
			givenDatastore: NewDatastore(),
			expected:       []model.Household{},
		},
		{
			msg: "not_empty",
			givenDatastore: &Datastore{
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
			t.Parallel()
			// When:
			result := tc.givenDatastore.RetrieveHouseholds()

			// Then:
			require.Equal(t, len(tc.expected), len(result))
			require.ElementsMatch(t, tc.expected, result)
		})
	}
}

func TestUpdateHousehold(t *testing.T) {
	testCases := []struct {
		msg            string
		givenDatastore *Datastore
		givenHousehold model.Household
		expectedError  error
	}{
		{
			msg: "success",
			givenDatastore: &Datastore{
				Households: map[int]model.Household{
					1: {
						ID:      1,
						Type:    "HDB",
						Members: []model.FamilyMember{},
					},
					2: {
						ID:   2,
						Type: "Landed",
						Members: []model.FamilyMember{
							{ID: 1, Name: "Jacky"},
						},
					},
				},
			},
			givenHousehold: model.Household{
				ID:   2,
				Type: "Landed",
				Members: []model.FamilyMember{
					{ID: 1, Name: "Jacky"},
					{ID: 2, Name: "Jeanny"},
				},
			},
			expectedError: nil,
		},
		{
			msg: "failure",
			givenDatastore: &Datastore{
				Households: map[int]model.Household{
					1: {
						ID:      1,
						Type:    "HDB",
						Members: []model.FamilyMember{},
					},
				},
			},
			givenHousehold: model.Household{
				Type:    "Condominium",
				Members: []model.FamilyMember{},
			},
			expectedError: ErrHouseholdInvalid,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			err := tc.givenDatastore.UpdateHousehold(tc.givenHousehold)

			// Then:
			require.Equal(t, tc.expectedError, err)

			actual := tc.givenDatastore.Households[tc.givenHousehold.ID]
			if tc.expectedError == nil {
				require.EqualValues(t, tc.givenHousehold, actual)
			}
		})
	}
}
