package corebussinesserrors

import "fmt"

// NotFoundError ...
type NotFoundError struct {
	status  uint16
	method  string
	message string
}

func (err *NotFoundError) Error() string {
	err.status = 404
	err.message = "Not Found Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
