package sessionframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	migrations "gomux_gorm/src/core/database/table_models"
)

type userRepository struct {
	db *gorm.DB
}

// IUserRepository ...
type IUserRepository interface {
	FindByEmail(email string) *migrations.Users
}

func (r *userRepository) FindByEmail(email string) *migrations.Users {
	user := migrations.Users{}

	r.db.First(&user, "email =?", email)

	return &user
}

// UserRepositoryConstructor ...
func UserRepositoryConstructor(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
