package core

import "fmt"

// UnauthorizedError ...
type UnauthorizedError struct {
	status uint16
	method string
}

func (err *UnauthorizedError) Error() string {
	err.status = 401

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
