package core

import "fmt"

// ForbiddenError ...
type ForbiddenError struct {
	status  uint16
	method  string
	message string
}

// Forbidden ...
func (err *ForbiddenError) Error() string {
	err.status = 403
	err.message = "Forbidden Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
