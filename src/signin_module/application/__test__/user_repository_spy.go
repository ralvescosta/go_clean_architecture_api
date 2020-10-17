package test

import (
	"github.com/stretchr/testify/mock"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
	entities "gomux_gorm/src/signin_module/bussiness/entities"
)

// UserRepositorySpy ...
type UserRepositorySpy struct {
	mock.Mock
}

// Create ...
func (u *UserRepositorySpy) Create(registerUser *entities.RegisterUsersEntity) *tables.Users {
	args := u.Called(registerUser)
	return args.Get(0).(*tables.Users)
}

// FindByEmail ...
func (u *UserRepositorySpy) FindByEmail(email string) *tables.Users {
	args := u.Called(email)
	return args.Get(0).(*tables.Users)
}
