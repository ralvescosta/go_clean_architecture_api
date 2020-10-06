package core

import "fmt"

// InternalServerError ...
type InternalServerError struct {
	status  uint16
	method  string
	message string
}

// InternalServer ...
func (err *InternalServerError) Error() string {
	err.status = 500
	err.message = "Internal Server Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
