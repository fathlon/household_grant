package handler

import (
	"errors"
	"testing"

	"github.com/fathlon/household_grant/model"
	"github.com/stretchr/testify/require"
)

func TestMapSearchOperation(t *testing.T) {
	testCases := []struct {
		msg            string
		givenParams    map[string][]string
		expectedError  error
		expectedResult model.SearchOperation
	}{
		{
			msg: "ignore_duplicate_values",
			givenParams: map[string][]string{
				"annual_income_gt": {"51000"},
				"household_size":   {"3", "6"},
				"has_couple":       {"true"},
			},
			expectedError: nil,
			expectedResult: model.SearchOperation{
				AnnualIncomeGT: 51000,
				HouseholdSize:  3,
				HasCouple:      true,
			},
		},
		{
			msg: "ignore_unknown_fields",
			givenParams: map[string][]string{
				"age_lt":         {"50"},
				"cant_find_me":   {"hi"},
				"household_size": {"6"},
			},
			expectedError: nil,
			expectedResult: model.SearchOperation{
				AgeLT:         50,
				HouseholdSize: 6,
			},
		},
		{
			msg: "wrong_field_value_int",
			givenParams: map[string][]string{
				"age_lt": {"wrong_type"},
			},
			expectedError:  errors.New("invalid type of value"),
			expectedResult: model.SearchOperation{},
		},
		{
			msg: "wrong_field_value_bool",
			givenParams: map[string][]string{
				"has_children": {"wrong_type"},
			},
			expectedError:  errors.New("invalid type of value"),
			expectedResult: model.SearchOperation{},
		},
		{
			msg: "skip_key_with_no_value",
			givenParams: map[string][]string{
				"household_income": {"230000"},
				"household_size":   {},
			},
			expectedError: nil,
			expectedResult: model.SearchOperation{
				HouseholdIncome: 230000,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			// When:
			result, err := mapSearchOperation(tc.givenParams)

			// Then:
			require.Equal(t, tc.expectedError, err)

			if tc.expectedError == nil {
				require.EqualValues(t, tc.expectedResult, result)
			}
		})
	}
}
