package db

import "errors"

var (
	// ErrHouseholdDuplicateID occurs when the index is already used
	ErrHouseholdDuplicateID = errors.New("duplicate household index")
	// ErrHouseholdNotFound occurs when household not found
	ErrHouseholdNotFound = errors.New("household not found")
	// ErrHouseholdInvalid occurs when household is invalid
	ErrHouseholdInvalid = errors.New("household invalid")

	// ErrFamilyMemberDuplicateID occurs when the index is already used
	ErrFamilyMemberDuplicateID = errors.New("duplicate family member index")
	// ErrFamilyMemberDuplicateName occurs when the name is duplicate
	ErrFamilyMemberDuplicateName = errors.New("duplicate name")
)
