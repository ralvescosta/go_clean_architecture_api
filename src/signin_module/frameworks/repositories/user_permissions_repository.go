package signinframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	migrations "gomux_gorm/src/core/database/table_models"
)

type usersPermissionsRepository struct {
	db *gorm.DB
}

// IUsersPermissionsRepository ...
type IUsersPermissionsRepository interface {
	Create(user *migrations.Users, permissionID int64, permission string) *migrations.UsersPermissions
}

func (r *usersPermissionsRepository) Create(user *migrations.Users, permissionID int64, permission string) *migrations.UsersPermissions {

	userPermission := migrations.UsersPermissions{
		PermissionID:   permissionID,
		PermissionRole: permission,
		UserID:         user.ID,
		UserName:       user.Name,
		UserEmail:      user.Email,
	}

	r.db.Create(&userPermission)

	return &userPermission
}

// UsersPermissionsRepositoryConstructor ...
func UsersPermissionsRepositoryConstructor(db *gorm.DB) IUsersPermissionsRepository {
	return &usersPermissionsRepository{db}
}
