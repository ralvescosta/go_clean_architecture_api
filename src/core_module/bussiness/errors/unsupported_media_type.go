package corebussinesserrors

import "fmt"

// UnsupportedMediaTypeError ...
type UnsupportedMediaTypeError struct {
	status  uint16
	method  string
	message string
}

func (err *UnsupportedMediaTypeError) Error() string {
	err.status = 415
	err.message = "Unsupported Media Type Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
