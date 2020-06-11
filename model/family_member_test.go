package model

import (
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidateFamilyMember(t *testing.T) {
	validDOB := time.Now().AddDate(-20, 0, 0)

	testCases := []struct {
		given    FamilyMember
		expected error
	}{
		{
			FamilyMember{
				Name:           "Jackie",
				Gender:         "M",
				OccupationType: "Unemployed",
				MaritalStatus:  "Married",
				DOB:            validDOB,
			},
			nil,
		},
		{
			FamilyMember{
				Name:           "Alexia",
				Gender:         "F",
				OccupationType: "Employed",
				MaritalStatus:  "Single",
				DOB:            validDOB,
			},
			nil,
		},
		{
			FamilyMember{
				Name:           "Alexia",
				Gender:         "F",
				OccupationType: "Student",
				MaritalStatus:  "Single",
				DOB:            validDOB,
			},
			nil,
		},
		{
			FamilyMember{
				Name: "",
			},
			ErrFamilyMemberNameInvalid,
		},
		{
			FamilyMember{
				Name:           "Jippi",
				Gender:         "T",
				OccupationType: "Employed",
				MaritalStatus:  "Married",
				DOB:            validDOB,
			},
			ErrFamilyMemberGenderInvalid,
		},
		{
			FamilyMember{
				Name:           "Jackie",
				Gender:         "F",
				OccupationType: "Freelancing",
				DOB:            validDOB,
			},
			ErrFamilyMemberOccupationTypeInvalid,
		},
		{
			FamilyMember{
				Name:           "Jackie",
				Gender:         "F",
				OccupationType: "Student",
				MaritalStatus:  "Complicated",
				DOB:            validDOB,
			},
			ErrFamilyMemberMaritalStatusInvalid,
		},
		{
			FamilyMember{
				Name:           "Jackie",
				Gender:         "F",
				OccupationType: "Student",
				MaritalStatus:  "Single",
			},
			ErrFamilyMemberDOBInvalid,
		},
		{
			FamilyMember{
				Name:           "Jackie",
				Gender:         "F",
				OccupationType: "Student",
				MaritalStatus:  "Single",
				DOB:            time.Now().AddDate(0, 0, 1),
			},
			ErrFamilyMemberDOBInvalid,
		},
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

func TestAge(t *testing.T) {
	now := time.Now()
	testCases := []struct {
		given    FamilyMember
		expected int
	}{
		{
			given: FamilyMember{
				DOB: now.AddDate(-25, 0, 0),
			},
			expected: 25,
		},
		{
			given: FamilyMember{
				DOB: now.AddDate(-25, 1, 0),
			},
			expected: 24,
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()
			// When:
			result := tc.given.Age()

			// Then:
			require.Equal(t, tc.expected, result)
		})
	}
}
