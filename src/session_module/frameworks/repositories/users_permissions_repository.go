package sessionframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

type userPermissionsRepository struct {
	db *gorm.DB
}

// IUsersPermissionsRepository ...
type IUsersPermissionsRepository interface {
	FindUserPermissions(userID int64) *[]tables.UsersPermissions
}

func (r *userPermissionsRepository) FindUserPermissions(userID int64) *[]tables.UsersPermissions {
	userPermissions := &[]tables.UsersPermissions{}

	r.db.Find(userPermissions, "id =?", userID)

	return userPermissions
}

// UsersPermissionsRepository ...
func UsersPermissionsRepository(db *gorm.DB) IUsersPermissionsRepository {
	return &userPermissionsRepository{db}
}
