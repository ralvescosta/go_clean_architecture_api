package core

// NotFoundError ...
type NotFoundError struct{}

func (*NotFoundError) Error() string {
	return "404 - Not Found Error"
}
