package signinframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
	entities "gomux_gorm/src/signin_module/bussiness/entities"
)

type userRepository struct {
	db *gorm.DB
}

// IUserRepository ...
type IUserRepository interface {
	Create(registerUser *entities.RegisterUsersEntity) *tables.Users
	FindByEmail(email string) *tables.Users
}

// Create ...
func (r *userRepository) Create(registerUser *entities.RegisterUsersEntity) *tables.Users {

	user := tables.Users{
		Name:     registerUser.Name,
		LastName: registerUser.LastName,
		Email:    registerUser.Email,
		Password: registerUser.Password,
	}

	r.db.Create(&user)

	return &user
}

// FindOne ...
func (r *userRepository) FindOne(id int) *tables.Users {
	user := tables.Users{}

	r.db.First(&user, id)

	return &user
}

// FindByEmail ...
func (r *userRepository) FindByEmail(email string) *tables.Users {
	user := tables.Users{}

	r.db.First(&user, "email =?", email)

	return &user
}

// FindAll ...
func (r *userRepository) FindAll() *[]tables.Users {

	user := []tables.Users{tables.Users{}}

	r.db.Find(&user)

	return &user
}

// Update ...
func (r *userRepository) Update(id int) *tables.Users {
	user := tables.Users{}

	return &user
}

// Delete ...
func (r *userRepository) Delete(id int) *tables.Users {
	user := tables.Users{}

	return &user
}

// UserRepository ...
func UserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
