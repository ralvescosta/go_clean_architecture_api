package signinapplicationusecases

import (
	core "gomux_gorm/src/core_module/bussiness/errors"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"
	entities "gomux_gorm/src/signin_module/bussiness/entities"
	crypto "gomux_gorm/src/signin_module/frameworks/crypto"
	repositories "gomux_gorm/src/signin_module/frameworks/repositories"

	"testing"
)

/*
* User Repository Spy
 */
type userRepositoryStructSpy struct{}

func (*userRepositoryStructSpy) Create(registerUser *entities.RegisterUsersEntity) *tables.Users {
	return &tables.Users{}
}
func (*userRepositoryStructSpy) FindByEmail(email string) *tables.Users {
	return &tables.Users{}
}
func UserRepositorySpy() repositories.IUserRepository {
	return &userRepositoryStructSpy{}
}

/*
* Permission Repository Spy
 */
type userPermissionsRepositoryStructSpy struct{}

func (*userPermissionsRepositoryStructSpy) Create(user *tables.Users, permissionID int64, permission string) *tables.UsersPermissions {
	return &tables.UsersPermissions{}
}
func UserPermissionsRepositorySpy() repositories.IUsersPermissionsRepository {
	return &userPermissionsRepositoryStructSpy{}
}

/*
* Crypto Spy
 */
type cryptoStructSpy struct{}

func (*cryptoStructSpy) HashPassword(password string) (string, error) {
	return "", nil
}
func CryptoSpy() crypto.IHasher {
	return &cryptoStructSpy{}
}

func TestSigninUsecase(t *testing.T) {
	userRepositorySpy := UserRepositorySpy()
	userPermissionsRepositorySpy := UserPermissionsRepositorySpy()
	cryptoSpy := CryptoSpy()

	sut := SigninUsecase(&userRepositorySpy, &userPermissionsRepositorySpy, &cryptoSpy)

	result := sut.SigninUsecase(&bussiness.RegisterUsersEntity{})

	switch result.(type) {
	case *core.UnauthorizedError:
		t.Errorf("Sum was incorrect")
		break
	}
}
