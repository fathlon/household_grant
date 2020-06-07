package service

// ValidationError represents error resulting from validation
type ValidationError struct {
	err error
}

// NewValidationError returns a new ValidationError
func NewValidationError(err error) *ValidationError {
	return &ValidationError{err}
}

func (e *ValidationError) Error() string {
	return e.err.Error()
}
