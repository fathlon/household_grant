package db

import "errors"

var (
	// ErrHouseholdDuplicateID occurs when the index is already used
	ErrHouseholdDuplicateID = errors.New("duplicate household index")
	// ErrHouseholdNotFound occurs when household not found
	ErrHouseholdNotFound = errors.New("household not found")

	// ErrFamilyMemberDuplicateID occurs when the index is already used
	ErrFamilyMemberDuplicateID = errors.New("duplicate family member index")
)
