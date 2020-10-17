package test

import (
	"github.com/stretchr/testify/mock"

	bussiness "gomux_gorm/src/auth_module/bussiness/entities"
	token "gomux_gorm/src/auth_module/frameworks/token"
)

// DecodedTokenSpy ...
type DecodedTokenSpy struct {
	mock.Mock
	token.IDecodedToken
}

// Decoded ...
func (d *DecodedTokenSpy) Decoded(t string) (*bussiness.TokenDecodedEntity, error) {
	args := d.Called(t)

	return args.Get(0).(*bussiness.TokenDecodedEntity), args.Error(1)
}
