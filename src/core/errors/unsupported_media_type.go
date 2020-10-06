package core

import "fmt"

// UnsupportedMediaTypeError ...
type UnsupportedMediaTypeError struct {
	status uint16
	method string
}

func (err *UnsupportedMediaTypeError) Error() string {
	err.status = 415

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
