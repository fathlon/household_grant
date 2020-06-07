package db

import "errors"

var (
	// ErrDuplicateID occurs when the index is already used
	ErrDuplicateID = errors.New("duplicate index")

	// ErrHouseholdTypeInvalid occurs when the household type is not within accepted options
	ErrHouseholdTypeInvalid = errors.New("invalid household type")
)
