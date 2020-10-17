package test

import "github.com/stretchr/testify/mock"

// CreateTokenSpy ...
type CreateTokenSpy struct {
	mock.Mock
}

// CreateToken ...
func (c *CreateTokenSpy) CreateToken(userID *int64, permissionID *int64) (string, error) {
	args := c.Called(userID, permissionID)
	return args.String(0), args.Error(1)
}
