package model

import "errors"

var (
	// ErrHouseholdTypeInvalid occurs when the household type is not within accepted options
	ErrHouseholdTypeInvalid = errors.New("invalid household type")
	// ErrHouseholdFamilyMemberNotExists occurs when the family member does not exists in the household
	ErrHouseholdFamilyMemberNotExists = errors.New("family member does not exists in household")

	// ErrFamilyMemberNameInvalid occurs when the name is invalid
	ErrFamilyMemberNameInvalid = errors.New("invalid name")
	// ErrFamilyMemberGenderInvalid occurs when the gender is not within accepted options
	ErrFamilyMemberGenderInvalid = errors.New("invalid gender")
	// ErrFamilyMemberOccupationTypeInvalid occurs when the occupation type is not within accepted options
	ErrFamilyMemberOccupationTypeInvalid = errors.New("invalid occupation type")
	// ErrFamilyMemberMaritalStatusInvalid occurs when the marital status is not within accepted options
	ErrFamilyMemberMaritalStatusInvalid = errors.New("invalid marital status")
	// ErrFamilyMemberDOBInvalid occurs when the DOB is invalid
	ErrFamilyMemberDOBInvalid = errors.New("invalid DOB")

	// ErrSearchConflictingComparison occurs when eq and (gt or lt) is both specified
	ErrSearchConflictingComparison = errors.New("conflicting comparison for search")
)
