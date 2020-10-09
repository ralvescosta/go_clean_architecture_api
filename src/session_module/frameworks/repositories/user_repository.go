package sessionframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

type userRepository struct {
	db *gorm.DB
}

// IUserRepository ...
type IUserRepository interface {
	FindByEmail(email string) *tables.Users
}

func (r *userRepository) FindByEmail(email string) *tables.Users {
	user := tables.Users{}

	r.db.First(&user, "email =?", email)

	return &user
}

// UserRepository ...
func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
