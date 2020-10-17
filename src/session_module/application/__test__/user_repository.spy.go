package test

import (
	"github.com/stretchr/testify/mock"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

// UserRepositorySpy ...
type UserRepositorySpy struct {
	mock.Mock
}

// FindByEmail ...
func (u *UserRepositorySpy) FindByEmail(email string) *tables.Users {
	args := u.Called(email)
	return args.Get(0).(*tables.Users)
}
