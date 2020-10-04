package signinframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	migrations "gomux_gorm/src/core/database/table_models"
	entities "gomux_gorm/src/signin_module/bussiness/entities"
)

type userRepository struct {
	db *gorm.DB
}

// IUserRepository ...
type IUserRepository interface {
	Create(registerUser *entities.RegisterUsersEntity) *migrations.Users
	FindByEmail(email string) *migrations.Users
}

// Create ...
func (r *userRepository) Create(registerUser *entities.RegisterUsersEntity) *migrations.Users {

	user := migrations.Users{
		Name:     registerUser.Name,
		LastName: registerUser.LastName,
		Email:    registerUser.Email,
		Password: registerUser.Password,
	}

	r.db.Create(&user)

	return &user
}

// FindOne ...
func (r *userRepository) FindOne(id int) *migrations.Users {
	user := migrations.Users{}

	r.db.First(&user, id)

	return &user
}

// FindByEmail ...
func (r *userRepository) FindByEmail(email string) *migrations.Users {
	user := migrations.Users{}

	r.db.First(&user, "email =?", email)

	return &user
}

// FindAll ...
func (r *userRepository) FindAll() *[]migrations.Users {

	user := []migrations.Users{migrations.Users{}}

	r.db.Find(&user)

	return &user
}

// Update ...
func (r *userRepository) Update(id int) *migrations.Users {
	user := migrations.Users{}

	return &user
}

// Delete ...
func (r *userRepository) Delete(id int) *migrations.Users {
	user := migrations.Users{}

	return &user
}

// UserRepositoryConstructor ...
func UserRepositoryConstructor(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
