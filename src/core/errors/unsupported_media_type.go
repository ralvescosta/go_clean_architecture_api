package core

// UnsupportedMediaTypeError ...
type UnsupportedMediaTypeError struct{}

func (*UnsupportedMediaTypeError) Error() string {
	return "415 - Unsupported Media Type Error"
}
