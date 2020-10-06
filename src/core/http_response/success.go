package core

// Success ...
type Success struct {
	status  uint16
	method  string
	message string
}

// Success ...
func (err *Success) Success() string {
	err.status = 200
	err.message = "Ok"

	return "Created"
}
