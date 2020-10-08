package sessionframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	tables "gomux_gorm/src/core/database/table_models"
)

type userRepository struct {
	db *gorm.DB
}

// IUserRepository ...
type IUserRepository interface {
	FindByID(id int64) *tables.Users
}

// FindByID ...
func (r *userRepository) FindByID(id int64) *tables.Users {
	user := tables.Users{}

	r.db.First(&user, "id =?", id)

	return &user
}

// UserRepository ...
func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
