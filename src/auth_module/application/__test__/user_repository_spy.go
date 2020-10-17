package test

import (
	"github.com/stretchr/testify/mock"

	repositories "gomux_gorm/src/auth_module/frameworks/repositories"
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

// UserRepositorySpy ...
type UserRepositorySpy struct {
	mock.Mock
	repositories.IUserRepository
}

// FindByID ...
func (u *UserRepositorySpy) FindByID(id int64) *tables.Users {
	args := u.Called(id)
	return args.Get(0).(*tables.Users)
}
