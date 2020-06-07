package handler

import (
	"net/http"

	"github.com/fathlon/household_grant/service"
)

// CheckError checks the given err and return as WebError
func CheckError(err error) (string, int) {
	code := http.StatusInternalServerError

	switch err.(type) {
	case *service.ValidationError:
		code = http.StatusBadRequest
	}

	return err.Error(), code
}
