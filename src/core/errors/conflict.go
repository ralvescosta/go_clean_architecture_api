package core

import "fmt"

// ConflictError ...
type ConflictError struct {
	status uint16
	method string
}

// ConflictError ...
func (err *ConflictError) Error() string {
	err.status = 409

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
