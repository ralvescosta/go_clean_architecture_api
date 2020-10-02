package signin_frameworks_repositories

import (
	migrations "gomux_gorm/src/core/database/table_models"
	entities "gomux_gorm/src/signin_module/bussiness/entities"

	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	Create(registerUser *entities.RegisterUsersEntity)
	FindByEmail(email string) *migrations.Users
}

func (r *userRepository) Create(registerUser *entities.RegisterUsersEntity) {

	r.db.Create(&migrations.Users{
		Name:     registerUser.Name,
		LastName: registerUser.LastName,
		Email:    registerUser.Email,
		Password: registerUser.Password,
	})
}

func (r *userRepository) FindOne(id int) *migrations.Users {
	user := migrations.Users{}

	r.db.First(&user, id)

	return &user
}

func (r *userRepository) FindByEmail(email string) *migrations.Users {
	user := migrations.Users{}

	r.db.First(&user, "email =?", email)

	return &user
}

func (r *userRepository) FindAll() *[]migrations.Users {

	user := []migrations.Users{migrations.Users{}}

	r.db.Find(&user)

	return &user
}

func (r *userRepository) Update(id int) *migrations.Users {
	user := migrations.Users{}

	return &user
}

func (r *userRepository) Delete(id int) *migrations.Users {
	user := migrations.Users{}

	return &user
}

func UserRepositoryConstructor(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}
