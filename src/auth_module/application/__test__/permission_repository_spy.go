package test

import (
	"github.com/stretchr/testify/mock"

	repositories "gomux_gorm/src/auth_module/frameworks/repositories"
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

// PermissionRepositorySpy ...
type PermissionRepositorySpy struct {
	mock.Mock
	repositories.IPermissionRepository
}

// FindByID ...
func (p *PermissionRepositorySpy) FindByID(id int64) *tables.Permissions {
	args := p.Called(id)

	return args.Get(0).(*tables.Permissions)
}
