package corebussinesserrors

import "fmt"

// ConflictError ...
type ConflictError struct {
	status  uint16
	method  string
	message string
}

// ConflictError ...
func (err *ConflictError) Error() string {
	err.status = 409
	err.message = "Conflict Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
