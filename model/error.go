package model

import "errors"

var (
	// ErrHouseholdTypeInvalid occurs when the household type is not within accepted options
	ErrHouseholdTypeInvalid = errors.New("invalid household type")
)
