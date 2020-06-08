package db

import "errors"

var (
	// ErrHouseholdDuplicateID occurs when the index is already used
	ErrHouseholdDuplicateID = errors.New("duplicate household index")
	// ErrFamilyMemberDuplicateID occurs when the index is already used
	ErrFamilyMemberDuplicateID = errors.New("duplicate family member index")
)
