package search

import (
	"testing"
	"time"

	"github.com/fathlon/household_grant/db"
	"github.com/fathlon/household_grant/model"
	"github.com/fathlon/household_grant/service"
	"github.com/stretchr/testify/require"
)

var (
	now = time.Now()
	hb1 = model.FamilyMember{
		Name:           "Husband1",
		Gender:         "M",
		MaritalStatus:  "Married",
		Spouse:         "Wife1",
		OccupationType: "Employed",
		AnnualIncome:   78000,
		DOB:            now.AddDate(-46, 0, 0),
	}
	wf1 = model.FamilyMember{
		Name:           "Wife1",
		Gender:         "F",
		MaritalStatus:  "Married",
		Spouse:         "Husband1",
		OccupationType: "Employed",
		AnnualIncome:   85000,
		DOB:            now.AddDate(-44, 0, 0),
	}
	ch1 = model.FamilyMember{
		Name:           "Child1",
		Gender:         "F",
		MaritalStatus:  "Single",
		OccupationType: "Student",
		DOB:            now.AddDate(-19, 0, 0),
	}
	ch2 = model.FamilyMember{
		Name:           "Child2",
		Gender:         "F",
		MaritalStatus:  "Single",
		OccupationType: "Student",
		DOB:            now.AddDate(-13, 0, 0),
	}
	eld1 = model.FamilyMember{
		Name:           "Elder1",
		Gender:         "M",
		MaritalStatus:  "Single",
		OccupationType: "Unemployed",
		DOB:            now.AddDate(-67, 0, 0),
	}
	sm1 = model.FamilyMember{
		Name:           "Single Mum",
		Gender:         "F",
		MaritalStatus:  "Single",
		OccupationType: "Employed",
		AnnualIncome:   250000,
		DOB:            now.AddDate(-34, 0, 0),
	}
	ch3 = model.FamilyMember{
		Name:   "Newborn",
		Gender: "M",
		DOB:    now.AddDate(0, -3, 0),
	}

	// couple with 2 children, total income over 163,000
	h1 = model.Household{
		ID:      1,
		Type:    "HDB",
		Members: []model.FamilyMember{hb1, wf1, ch1, ch2},
	}
	// single elder no income
	h2 = model.Household{
		ID:      2,
		Type:    "HDB",
		Members: []model.FamilyMember{eld1},
	}
	// single mother with newborn, total income over 250000
	h3 = model.Household{
		ID:      3,
		Type:    "Landed",
		Members: []model.FamilyMember{sm1, ch3},
	}
)

func TestRetrieves(t *testing.T) {
	testCases := []struct {
		msg                  string
		givenSearchOperation model.SearchOperation
		expectedError        error
		expected             []model.Household
	}{
		{
			msg: "validation_error",
			givenSearchOperation: model.SearchOperation{
				Age:   5,
				AgeGT: 6,
			},
			expectedError: service.NewValidationError(model.ErrSearchConflictingComparison),
			expected:      []model.Household{},
		},
		{
			msg: "success",
			givenSearchOperation: model.SearchOperation{
				AgeGT: 50,
			},
			expectedError: nil,
			expected:      []model.Household{h2},
		},
		{
			msg: "household_only",
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeGT: 1,
				AnnualIncomeGT:  80000,
				WholeHousehold:  true,
			},
			expectedError: nil,
			expected: []model.Household{
				{
					ID:      1,
					Type:    "HDB",
					Members: []model.FamilyMember{},
				},
				{
					ID:      3,
					Type:    "Landed",
					Members: []model.FamilyMember{},
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// Given:
			ds := &db.Datastore{
				Households: map[int]model.Household{
					1: h1,
					2: h2,
					3: h3,
				},
			}

			// When:
			result, err := Retrieves(ds, tc.givenSearchOperation)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.ElementsMatch(t, tc.expected, result)
			}
		})
	}
}
