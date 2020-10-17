package test

import (
	"github.com/stretchr/testify/mock"
)

// CryptoSpy ...
type CryptoSpy struct {
	mock.Mock
}

// HashPassword ...
func (c *CryptoSpy) HashPassword(password string) (string, error) {
	args := c.Called(password)
	return args.String(0), args.Error(1)
}
