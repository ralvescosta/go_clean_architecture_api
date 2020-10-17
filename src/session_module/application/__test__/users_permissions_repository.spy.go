package test

import (
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"

	"github.com/stretchr/testify/mock"
)

// UsersPermissionsRepositorySpy ...
type UsersPermissionsRepositorySpy struct {
	mock.Mock
}

// FindUserPermissions ...
func (u *UsersPermissionsRepositorySpy) FindUserPermissions(userID int64) *[]tables.UsersPermissions {
	args := u.Called(userID)
	return args.Get(0).(*[]tables.UsersPermissions)
}
