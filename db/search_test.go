package db

import (
	"strconv"
	"testing"
	"time"

	"github.com/fathlon/household_grant/model"
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
	// household with no members
	h4 = model.Household{
		ID:      4,
		Type:    "HDB",
		Members: []model.FamilyMember{},
	}
)

func TestSearch(t *testing.T) {
	testCases := []struct {
		givenSearchOperation model.SearchOperation
		expected             []model.Household
	}{
		{
			givenSearchOperation: model.SearchOperation{
				HouseholdIncomeLT: 200000,
				AgeLT:             20,
			},
			expected: []model.Household{
				{
					ID:      1,
					Type:    "HDB",
					Members: []model.FamilyMember{ch1, ch2},
				},
			},
		},
		{
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeLT: 4,
			},
			expected: []model.Household{h2, h3},
		},
		{
			givenSearchOperation: model.SearchOperation{
				AnnualIncomeLT: 80000,
			},
			expected: []model.Household{
				{
					ID:      1,
					Type:    "HDB",
					Members: []model.FamilyMember{hb1, ch1, ch2},
				},
				{
					ID:      2,
					Type:    "HDB",
					Members: []model.FamilyMember{eld1},
				},
				{
					ID:      3,
					Type:    "Landed",
					Members: []model.FamilyMember{ch3},
				},
			},
		},
		{
			givenSearchOperation: model.SearchOperation{},
			expected:             []model.Household{},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// Given:
			ds := &Datastore{
				Households: map[int]model.Household{
					1: h1,
					2: h2,
					3: h3,
				},
			}

			// When:
			result := ds.Search(tc.givenSearchOperation)

			// Then:
			require.ElementsMatch(t, tc.expected, result)
		})
	}
}

func TestMatchHouseholdIncome(t *testing.T) {
	testCases := []struct {
		msg                  string
		givenSearchOperation model.SearchOperation
		givenTotalIncome     int
		expected             bool
	}{
		{
			msg: "equal_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncome: 15000,
			},
			givenTotalIncome: 15000,
			expected:         true,
		},
		{
			msg: "equal_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncome: 15000,
			},
			givenTotalIncome: 15001,
			expected:         false,
		},
		{
			msg: "gt_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncomeGT: 15000,
			},
			givenTotalIncome: 15001,
			expected:         true,
		},
		{
			msg: "gt_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncomeGT: 15000,
			},
			givenTotalIncome: 15000,
			expected:         false,
		},
		{
			msg: "lt_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncomeLT: 15000,
			},
			givenTotalIncome: 14999,
			expected:         true,
		},
		{
			msg: "lt_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncomeLT: 15000,
			},
			givenTotalIncome: 15000,
			expected:         false,
		},
		{
			msg: "gt_lt_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncomeLT: 15000,
				HouseholdIncomeGT: 14000,
			},
			givenTotalIncome: 14999,
			expected:         true,
		},
		{
			msg: "gt_lt_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdIncomeLT: 15000,
				HouseholdIncomeGT: 14000,
			},
			givenTotalIncome: 13000,
			expected:         false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result := matchHouseholdIncome(tc.givenTotalIncome, tc.givenSearchOperation)

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMatchHouseholdSize(t *testing.T) {
	testCases := []struct {
		msg                  string
		givenSearchOperation model.SearchOperation
		givenSize            int
		expected             bool
	}{
		{
			msg: "equal_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdSize: 3,
			},
			givenSize: 3,
			expected:  true,
		},
		{
			msg: "equal_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdSize: 3,
			},
			givenSize: 4,
			expected:  false,
		},
		{
			msg: "gt_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeGT: 3,
			},
			givenSize: 4,
			expected:  true,
		},
		{
			msg: "gt_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeGT: 3,
			},
			givenSize: 3,
			expected:  false,
		},
		{
			msg: "lt_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeLT: 4,
			},
			givenSize: 3,
			expected:  true,
		},
		{
			msg: "lt_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeLT: 4,
			},
			givenSize: 6,
			expected:  false,
		},
		{
			msg: "gt_lt_true",
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeLT: 6,
				HouseholdSizeGT: 2,
			},
			givenSize: 4,
			expected:  true,
		},
		{
			msg: "gt_lt_false",
			givenSearchOperation: model.SearchOperation{
				HouseholdSizeLT: 6,
				HouseholdSizeGT: 2,
			},
			givenSize: 1,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result := matchHouseholdSize(tc.givenSize, tc.givenSearchOperation)

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMatchAnnualIncome(t *testing.T) {
	testCases := []struct {
		msg                  string
		givenSearchOperation model.SearchOperation
		givenAnnualIncome    int
		expected             bool
	}{
		{
			msg: "equal_true",
			givenSearchOperation: model.SearchOperation{
				AnnualIncome: 15000,
			},
			givenAnnualIncome: 15000,
			expected:          true,
		},
		{
			msg: "equal_false",
			givenSearchOperation: model.SearchOperation{
				AnnualIncome: 15000,
			},
			givenAnnualIncome: 15001,
			expected:          false,
		},
		{
			msg: "gt_true",
			givenSearchOperation: model.SearchOperation{
				AnnualIncomeGT: 15000,
			},
			givenAnnualIncome: 15001,
			expected:          true,
		},
		{
			msg: "gt_false",
			givenSearchOperation: model.SearchOperation{
				AnnualIncomeGT: 15000,
			},
			givenAnnualIncome: 15000,
			expected:          false,
		},
		{
			msg: "lt_true",
			givenSearchOperation: model.SearchOperation{
				AnnualIncomeLT: 15000,
			},
			givenAnnualIncome: 14999,
			expected:          true,
		},
		{
			msg: "lt_false",
			givenSearchOperation: model.SearchOperation{
				AnnualIncomeLT: 15000,
			},
			givenAnnualIncome: 15000,
			expected:          false,
		},
		{
			msg: "gt_lt_true",
			givenSearchOperation: model.SearchOperation{
				AnnualIncomeLT: 15000,
				AnnualIncomeGT: 14000,
			},
			givenAnnualIncome: 14999,
			expected:          true,
		},
		{
			msg: "gt_lt_false",
			givenSearchOperation: model.SearchOperation{
				AnnualIncomeLT: 15000,
				AnnualIncomeGT: 14000,
			},
			givenAnnualIncome: 13000,
			expected:          false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result := matchAnnualIncome(tc.givenAnnualIncome, tc.givenSearchOperation)

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMatchAge(t *testing.T) {
	testCases := []struct {
		msg                  string
		givenSearchOperation model.SearchOperation
		givenSize            int
		expected             bool
	}{
		{
			msg: "equal_true",
			givenSearchOperation: model.SearchOperation{
				Age: 35,
			},
			givenSize: 35,
			expected:  true,
		},
		{
			msg: "equal_false",
			givenSearchOperation: model.SearchOperation{
				Age: 35,
			},
			givenSize: 41,
			expected:  false,
		},
		{
			msg: "gt_true",
			givenSearchOperation: model.SearchOperation{
				AgeGT: 30,
			},
			givenSize: 44,
			expected:  true,
		},
		{
			msg: "gt_false",
			givenSearchOperation: model.SearchOperation{
				AgeGT: 35,
			},
			givenSize: 31,
			expected:  false,
		},
		{
			msg: "lt_true",
			givenSearchOperation: model.SearchOperation{
				AgeLT: 47,
			},
			givenSize: 34,
			expected:  true,
		},
		{
			msg: "lt_false",
			givenSearchOperation: model.SearchOperation{
				AgeLT: 43,
			},
			givenSize: 64,
			expected:  false,
		},
		{
			msg: "gt_lt_true",
			givenSearchOperation: model.SearchOperation{
				AgeLT: 60,
				AgeGT: 25,
			},
			givenSize: 40,
			expected:  true,
		},
		{
			msg: "gt_lt_false",
			givenSearchOperation: model.SearchOperation{
				AgeLT: 60,
				AgeGT: 20,
			},
			givenSize: 20,
			expected:  false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result := matchAge(tc.givenSize, tc.givenSearchOperation)

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}
