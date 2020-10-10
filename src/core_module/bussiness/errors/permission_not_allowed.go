package corebussinesserrors

import "fmt"

// PermissionNotAllowedError ...
type PermissionNotAllowedError struct {
	status  uint16
	method  string
	message string
}

func (err *PermissionNotAllowedError) Error() string {
	err.status = 401
	err.message = "Permission Not Allowed Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
