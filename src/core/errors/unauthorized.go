package core

// UnauthorizedError ...
type UnauthorizedError struct{}

func (*UnauthorizedError) Error() string {
	return "401 - Unauthorized Error"
}
