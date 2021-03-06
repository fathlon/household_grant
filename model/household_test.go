package model

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateHousehold(t *testing.T) {
	testCases := []struct {
		given    Household
		expected error
	}{
		{Household{Type: "HDB"}, nil},
		{Household{Type: "Landed"}, nil},
		{Household{Type: "Condominium"}, nil},
		{Household{Type: "Special"}, ErrHouseholdTypeInvalid},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			// When:
			result := tc.given.Validate()

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestAddMember(t *testing.T) {
	testCases := []struct {
		msg            string
		givenHousehold Household
		givenMember    FamilyMember
		expectedCount  int
	}{
		{
			msg: "add_existing_member",
			givenHousehold: Household{
				Members: []FamilyMember{
					{ID: 1, Name: "Rachael"},
					{ID: 2, Name: "Monica"},
				},
			},
			givenMember:   FamilyMember{ID: 2, Name: "Monica"},
			expectedCount: 2,
		},
		{
			msg: "add_new_member",
			givenHousehold: Household{
				Members: []FamilyMember{
					{ID: 1, Name: "Rachael"},
					{ID: 2, Name: "Monica"},
				},
			},
			givenMember:   FamilyMember{ID: 3, Name: "Ross"},
			expectedCount: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			tc.givenHousehold.AddMember(tc.givenMember)

			// Then:
			require.Len(t, tc.givenHousehold.Members, tc.expectedCount)
		})
	}
}

func TestMemberExists(t *testing.T) {
	testCases := []struct {
		msg            string
		givenHousehold Household
		givenMemberID  int
		expected       bool
	}{
		{
			msg: "true",
			givenHousehold: Household{
				Members: []FamilyMember{
					{ID: 1, Name: "Rachael"},
					{ID: 2, Name: "Monica"},
				},
			},
			givenMemberID: 2,
			expected:      true,
		},
		{
			msg: "false",
			givenHousehold: Household{
				Members: []FamilyMember{
					{ID: 1, Name: "Rachael"},
				},
			},
			givenMemberID: 2,
			expected:      false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result := tc.givenHousehold.MemberExists(tc.givenMemberID)

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestDeleteMember(t *testing.T) {
	testCases := []struct {
		msg            string
		givenHousehold Household
		givenMemberID  int
		expected       bool
	}{
		{
			msg: "success",
			givenHousehold: Household{
				Members: []FamilyMember{
					{ID: 1, Name: "Rachael"},
					{ID: 2, Name: "Monica"},
				},
			},
			givenMemberID: 2,
			expected:      true,
		},
		{
			msg: "not_found",
			givenHousehold: Household{
				Members: []FamilyMember{},
			},
			givenMemberID: 2,
			expected:      false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.msg, func(t *testing.T) {
			t.Parallel()
			// When:
			result := tc.givenHousehold.DeleteMember(tc.givenMemberID)

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}
