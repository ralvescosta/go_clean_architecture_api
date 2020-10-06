package core

// BadRequestError ...
type BadRequestError struct{}

// BadRequest ...
func (*BadRequestError) Error() string {
	return "400 - Bad Request Error"
}
