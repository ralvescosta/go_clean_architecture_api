package sessionframeworksrepositories

import (
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"

	"github.com/jinzhu/gorm"
)

type permissionRepository struct {
	db *gorm.DB
}

// IPermissionRepository ...
type IPermissionRepository interface {
	FindByID(id int64) *tables.Permissions
}

// FindById ...
func (r *permissionRepository) FindByID(id int64) *tables.Permissions {
	permission := &tables.Permissions{}

	r.db.First(permission, "id =?", id)

	return permission
}

// PermissionRepository ...
func PermissionRepository(db *gorm.DB) IPermissionRepository {
	return &permissionRepository{db}
}
