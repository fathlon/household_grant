package db

import "errors"

var (
	// ErrDuplicateID occurs when the index is already used
	ErrDuplicateID = errors.New("duplicate index")
)
