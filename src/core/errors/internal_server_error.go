package core

// InternalServerError ...
type InternalServerError struct{}

// InternalServer ...
func (*InternalServerError) Error() string {
	return "500 - Internal Sever Error"
}
