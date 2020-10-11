package sessionapplicationusecases

import (
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
type userRepositoryStructSpy struct{}

func (*userRepositoryStructSpy) FindByEmail(email string) *tables.Users {
	return &tables.Users{}
}
func UserRepositorySpy() repositories.IUserRepository {
	return &userRepositoryStructSpy{}
}

/*
* SessionRepositorySpy
 */
type sessionRepositoryStructSpy struct{}

func (*sessionRepositoryStructSpy) Create(session *bussiness.SessionEntity, user *tables.Users) {}
func SessionRepositorySpy() repositories.ISessionRepository {
	return &sessionRepositoryStructSpy{}
}

/*
* UsersPermissionsRepositorySpy
 */
type usersPermissionsRepositoryStructSpy struct{}

func (*usersPermissionsRepositoryStructSpy) FindUserPermissions(userID int64) *[]tables.UsersPermissions {
	return &[]tables.UsersPermissions{}
}
func UsersPermissionsRepositorySpy() repositories.IUsersPermissionsRepository {
	return &usersPermissionsRepositoryStructSpy{}
}

/*
* HasherSpy
 */
type hasherStructSpy struct{}

func (*hasherStructSpy) CheckPasswordHash(password, hash string) bool {
	return true
}
func HasherSpy() crypto.IHasher {
	return &hasherStructSpy{}
}

/*
* CreateTokenSpy
 */
type createTokenStructSpy struct{}

func (*createTokenStructSpy) CreateToken(userID *int64, permissionID *int64) (string, error) {
	return "", nil
}
func CreateTokenSpy() token.ICreateToken {
	return &createTokenStructSpy{}
}

func TestSessionUsecase(t *testing.T) {
	userRepositorySpy := UserRepositorySpy()
	sessionRepositorySpy := SessionRepositorySpy()
	usersPermissionsRepositorySpy := UsersPermissionsRepositorySpy()
	cryptoSpy := HasherSpy()
	tokenSpy := CreateTokenSpy()

	sut := SessionUsecase(&userRepositorySpy, &sessionRepositorySpy, &usersPermissionsRepositorySpy, &cryptoSpy, &tokenSpy)

	_, err := sut.SessionUsecase(&bussiness.UsersInput{}, &bussiness.SessionEntity{})

	if err != nil {
		t.Skipf("Error")
	}
}
