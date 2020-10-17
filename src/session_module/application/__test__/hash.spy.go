package test

import "github.com/stretchr/testify/mock"

// HasherSpy ...
type HasherSpy struct {
	mock.Mock
}

// CheckPasswordHash ...
func (h *HasherSpy) CheckPasswordHash(password, hash string) bool {
	args := h.Called(password, hash)
	return args.Bool(0)
}
