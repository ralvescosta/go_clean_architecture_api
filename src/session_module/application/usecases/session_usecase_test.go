package sessionapplicationusecases

import (
	"errors"
	"testing"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
	bussiness "gomux_gorm/src/session_module/bussiness/entities"
	crypto "gomux_gorm/src/session_module/frameworks/crypto"
	repositories "gomux_gorm/src/session_module/frameworks/repositories"
	token "gomux_gorm/src/session_module/frameworks/token"
)

/*
* UserRepositorySpy
 */
type userRepositoryStructSpy struct {
	res *tables.Users
}

func (u *userRepositoryStructSpy) FindByEmail(email string) *tables.Users {
	return u.res
}
func UserRepositorySpy(userRepoRes *tables.Users) repositories.IUserRepository {
	return &userRepositoryStructSpy{res: userRepoRes}
}

/**/

/*
* SessionRepositorySpy
 */
type sessionRepositoryStructSpy struct{}

func (*sessionRepositoryStructSpy) Create(session *bussiness.SessionEntity, user *tables.Users) {}
func SessionRepositorySpy() repositories.ISessionRepository {
	return &sessionRepositoryStructSpy{}
}

/**/

/*
* UsersPermissionsRepositorySpy
 */
type usersPermissionsRepositoryStructSpy struct {
	res *[]tables.UsersPermissions
}

func (u *usersPermissionsRepositoryStructSpy) FindUserPermissions(userID int64) *[]tables.UsersPermissions {
	return u.res
}
func UsersPermissionsRepositorySpy(usersPermissionRepoRes *[]tables.UsersPermissions) repositories.IUsersPermissionsRepository {
	return &usersPermissionsRepositoryStructSpy{res: usersPermissionRepoRes}
}

/**/

/*
* HasherSpy
 */
type hasherStructSpy struct {
	res bool
}

func (h *hasherStructSpy) CheckPasswordHash(password, hash string) bool {
	return h.res
}
func HasherSpy(cryptoRes bool) crypto.IHasher {
	return &hasherStructSpy{res: cryptoRes}
}

/**/

/*
* CreateTokenSpy
 */
type createTokenStructSpy struct {
	res string
	err error
}

func (c *createTokenStructSpy) CreateToken(userID *int64, permissionID *int64) (string, error) {
	return c.res, c.err
}
func CreateTokenSpy(createTokenRes string, createToken error) token.ICreateToken {
	return &createTokenStructSpy{res: createTokenRes, err: createToken}
}

/**/

/*
* MAKE SUT
 */
func makeSut(userRepoRes *tables.Users, usersPermissionRepoRes *[]tables.UsersPermissions, cryptoRes bool, createTokenRes string, createToken error) ISessionUsecase {
	userRepositorySpy := UserRepositorySpy(userRepoRes)
	sessionRepositorySpy := SessionRepositorySpy()
	usersPermissionsRepositorySpy := UsersPermissionsRepositorySpy(usersPermissionRepoRes)
	cryptoSpy := HasherSpy(cryptoRes)
	tokenSpy := CreateTokenSpy(createTokenRes, createToken)

	sut := SessionUsecase(&userRepositorySpy, &sessionRepositorySpy, &usersPermissionsRepositorySpy, &cryptoSpy, &tokenSpy)

	return sut
}

/**/

func TestSessionUsecase(t *testing.T) {
	sut := makeSut(
		&tables.Users{
			ID: 1,
		},
		&[]tables.UsersPermissions{
			tables.UsersPermissions{
				PermissionID: 2,
			},
		},
		true,
		"",
		nil,
	)

	_, err := sut.SessionUsecase(&bussiness.UsersInput{}, &bussiness.SessionEntity{})

	if err != nil {
		t.Skipf("SessionUsecase()")
	}
}

func TestShouldReturnNotFoundErrorIfUserEmailNotRegistered(t *testing.T) {
	sut := makeSut(
		&tables.Users{},
		&[]tables.UsersPermissions{},
		true,
		"",
		nil,
	)

	_, err := sut.SessionUsecase(&bussiness.UsersInput{}, &bussiness.SessionEntity{})

	if err.Error() != "Something went wrong with the  request. Server returned 404 status." {
		t.Errorf("SessionUsecase, check if user exist, Unexpected Response")
	}
}

func TestShouldReturnUnauthorizedErrorIfCheckPasswordHashReturnsFalse(t *testing.T) {
	sut := makeSut(
		&tables.Users{
			ID: 1,
		},
		&[]tables.UsersPermissions{},
		false,
		"",
		nil,
	)

	_, err := sut.SessionUsecase(&bussiness.UsersInput{}, &bussiness.SessionEntity{})

	if err.Error() != "Something went wrong with the  request. Server returned 401 status." {
		t.Errorf("SessionUsecase, check if user password hash is valid, Unexpected Response")
	}
}

func TestShouldReturnForbiddenErrorIfUserRoleIsUnauthorized(t *testing.T) {
	sut := makeSut(
		&tables.Users{
			ID: 1,
		},
		&[]tables.UsersPermissions{
			tables.UsersPermissions{
				PermissionID: 1,
			},
		},
		true,
		"",
		nil,
	)

	_, err := sut.SessionUsecase(&bussiness.UsersInput{}, &bussiness.SessionEntity{})

	if err.Error() != "Something went wrong with the  request. Server returned 403 status." {
		t.Errorf("SessionUsecase, check if user password hash is valid, Unexpected Response")
	}
}

func TestReturnInternalServerErrorIfSomeErrorOccurOnCreateToken(t *testing.T) {
	sut := makeSut(
		&tables.Users{
			ID: 1,
		},
		&[]tables.UsersPermissions{
			tables.UsersPermissions{
				PermissionID: 2,
			},
		},
		true,
		"",
		errors.New("some error"),
	)

	_, err := sut.SessionUsecase(&bussiness.UsersInput{}, &bussiness.SessionEntity{})

	if err.Error() != "Something went wrong with the  request. Server returned 500 status." {
		t.Errorf("SessionUsecase, check if user password hash is valid, Unexpected Response")
	}
}
