package corebussinesserrors

import "fmt"

// TokenExpiredError ...
type TokenExpiredError struct {
	status  uint16
	method  string
	message string
}

func (err *TokenExpiredError) Error() string {
	err.status = 401
	err.message = "Token Expired Error"

	return fmt.Sprintf("Something went wrong with the %v request. Server returned %v status.",
		err.method, err.status)
}
