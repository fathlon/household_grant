package model

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidate(t *testing.T) {
	truePointer := true
	testCases := []struct {
		msg      string
		given    SearchOperation
		expected error
	}{
		{
			msg: "conflicting_household_income",
			given: SearchOperation{
				HouseholdIncome:   15000,
				HouseholdIncomeLT: 3500,
			},
			expected: ErrSearchConflictingComparison,
		},
		{
			msg: "conflicting_household_size",
			given: SearchOperation{
				HouseholdSize:   4,
				HouseholdSizeGT: 3,
			},
			expected: ErrSearchConflictingComparison,
		},
		{
			msg: "conflicting_age",
			given: SearchOperation{
				Age:   45,
				AgeLT: 30,
			},
			expected: ErrSearchConflictingComparison,
		},
		{
			msg: "conflicting_annual_income",
			given: SearchOperation{
				AnnualIncome:   300000,
				AnnualIncomeGT: 100000,
			},
			expected: ErrSearchConflictingComparison,
		},
		{
			msg: "success",
			given: SearchOperation{
				Age:             45,
				AnnualIncomeGT:  100000,
				HouseholdSizeGT: 3,
				HouseholdIncome: 150000,
				HasCouple:       &truePointer,
			},
			expected: nil,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			err := tc.given.Validate()

			// Then:
			require.Equal(t, tc.expected, err)
		})
	}
}

func TestSetFlags(t *testing.T) {
	testCases := []struct {
		given    SearchOperation
		expected SearchOperation
	}{
		{
			given: SearchOperation{
				HouseholdIncome: 200000,
				AgeGT:           16,
			},
			expected: SearchOperation{
				CompareHouseholdIncome: true,
				CompareAge:             true,
			},
		},
		{
			given: SearchOperation{
				AnnualIncomeLT:  100000,
				HouseholdSizeLT: 5,
			},
			expected: SearchOperation{
				CompareAnnualIncome:  true,
				CompareHouseholdSize: true,
			},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// When:
			tc.given.SetFlags()

			// Then:
			require.Equal(t, tc.expected.CompareHouseholdIncome, tc.given.CompareHouseholdIncome)
			require.Equal(t, tc.expected.CompareHouseholdSize, tc.given.CompareHouseholdSize)
			require.Equal(t, tc.expected.CompareAnnualIncome, tc.given.CompareAnnualIncome)
			require.Equal(t, tc.expected.CompareAge, tc.given.CompareAge)
		})
	}
}
