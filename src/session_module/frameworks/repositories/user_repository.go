package session_frameworks_repositories

import (
	migrations "gomux_gorm/src/core/database/table_models"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	FindByEmail(email string) *migrations.Users
}

func (r *userRepository) FindByEmail(email string) *migrations.Users {
	user := migrations.Users{}

	r.db.First(&user, "email =?", email)

	return &user
}

func UserRepositoryConstructor(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
