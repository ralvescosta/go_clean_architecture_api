package core

import "fmt"

// BadRequestError ...
type BadRequestError struct {
	status uint16
	method string
}

// BadRequest ...
func (err *BadRequestError) Error() string {
	err.status = 400

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
