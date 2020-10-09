package corebussinesserrors

import "fmt"

// UnauthorizedError ...
type UnauthorizedError struct {
	status  uint16
	method  string
	message string
}

func (err *UnauthorizedError) Error() string {
	err.status = 401
	err.message = "Unauthorized Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
