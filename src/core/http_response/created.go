package core

// Created ...
type Created struct {
	status  uint16
	method  string
	message string
}

// Success ...
func (err *Created) Success() string {
	err.status = 201
	err.message = "Ok"

	return "Created"
}
