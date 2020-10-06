package core

// NoBody ...
type NoBody struct {
	status  uint16
	method  string
	message string
}

// Success ...
func (err *NoBody) Success() string {
	err.status = 204
	err.message = "Ok"

	return "Created"
}
