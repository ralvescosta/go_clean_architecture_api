package signinframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

type usersPermissionsRepository struct {
	db *gorm.DB
}

// IUsersPermissionsRepository ...
type IUsersPermissionsRepository interface {
	Create(user *tables.Users, permissionID int64, permission string) *tables.UsersPermissions
}

func (r *usersPermissionsRepository) Create(user *tables.Users, permissionID int64, permission string) *tables.UsersPermissions {

	userPermission := tables.UsersPermissions{
		PermissionID:   permissionID,
		PermissionRole: permission,
		UserID:         user.ID,
		UserName:       user.Name,
		UserEmail:      user.Email,
	}

	r.db.Create(&userPermission)

	return &userPermission
}

// UsersPermissionsRepository ...
func UsersPermissionsRepository(db *gorm.DB) IUsersPermissionsRepository {
	return &usersPermissionsRepository{db}
}
