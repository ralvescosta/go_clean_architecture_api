package test

import (
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"

	"github.com/stretchr/testify/mock"
)

// UserPermissionsRepositorySpy ...
type UserPermissionsRepositorySpy struct {
	mock.Mock
}

// Create ...
func (u *UserPermissionsRepositorySpy) Create(user *tables.Users, permissionID int64, permission string) *tables.UsersPermissions {
	args := u.Called(user, permissionID, permission)
	return args.Get(0).(*tables.UsersPermissions)
}
