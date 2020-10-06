package core

import "fmt"

// BadRequestError ...
type BadRequestError struct {
	status  uint16
	method  string
	message string
}

// BadRequest ...
func (err *BadRequestError) Error() string {
	err.status = 400
	err.message = "Bad Request Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
