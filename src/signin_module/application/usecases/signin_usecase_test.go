package signinapplicationusecases

import (
	"errors"
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
type userRepositoryStructSpy struct {
	res *tables.Users
}

func (u *userRepositoryStructSpy) Create(registerUser *entities.RegisterUsersEntity) *tables.Users {
	return u.res
}
func (u *userRepositoryStructSpy) FindByEmail(email string) *tables.Users {
	return u.res
}
func UserRepositorySpy(userRepoRes *tables.Users) repositories.IUserRepository {
	return &userRepositoryStructSpy{res: userRepoRes}
}

/**/

/*
* Permission Repository Spy
 */
type userPermissionsRepositoryStructSpy struct {
	res *tables.UsersPermissions
}

func (u *userPermissionsRepositoryStructSpy) Create(user *tables.Users, permissionID int64, permission string) *tables.UsersPermissions {
	return u.res
}
func UserPermissionsRepositorySpy(userPermissionRepoRes *tables.UsersPermissions) repositories.IUsersPermissionsRepository {
	return &userPermissionsRepositoryStructSpy{res: userPermissionRepoRes}
}

/**/

/*
* Crypto Spy
 */
type cryptoStructSpy struct {
	res string
	err error
}

func (c *cryptoStructSpy) HashPassword(password string) (string, error) {
	return c.res, c.err
}
func CryptoSpy(cryptoRes string, cryptoErr error) crypto.IHasher {
	return &cryptoStructSpy{res: cryptoRes, err: cryptoErr}
}

/**/

/*
* MAKE SUT
 */
func makeSut(userRepoRes *tables.Users, userPermissionRepoRes *tables.UsersPermissions, cryptoRes string, cryptoErr error) ISigninUsecase {
	userRepositorySpy := UserRepositorySpy(userRepoRes)
	userPermissionsRepositorySpy := UserPermissionsRepositorySpy(userPermissionRepoRes)
	cryptoSpy := CryptoSpy(cryptoRes, cryptoErr)

	sut := SigninUsecase(&userRepositorySpy, &userPermissionsRepositorySpy, &cryptoSpy)

	return sut
}

/**/

func TestSigninUsecase(t *testing.T) {
	sut := makeSut(&tables.Users{}, &tables.UsersPermissions{}, "", nil)

	result := sut.SigninUsecase(&bussiness.RegisterUsersEntity{})

	if result != nil {
		t.Errorf("SigninUsecase()")
	}
}

func TestShouldReturnConflictErrorIfUserEmailAlreadyExist(t *testing.T) {
	sut := makeSut(
		&tables.Users{
			ID: 1,
		},
		&tables.UsersPermissions{},
		"",
		nil,
	)

	result := sut.SigninUsecase(&bussiness.RegisterUsersEntity{})

	if result.Error() != "Something went wrong with the  request. Server returned 409 status." {
		t.Errorf("SigninUsecase, check if user already exist, Unexpected Response")
	}
}

func TestShouldReturnUnauthorizedErrorIfSomeErrorOccurInCreateHashPassword(t *testing.T) {
	sut := makeSut(&tables.Users{}, &tables.UsersPermissions{}, "", errors.New("some error"))

	result := sut.SigninUsecase(&bussiness.RegisterUsersEntity{})

	if result.Error() != "Something went wrong with the  request. Server returned 401 status." {
		t.Errorf("SigninUsecase, check if created hash password, Unexpected Response")
	}
}
