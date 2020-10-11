package authapplicationusecases

import (
	"testing"

	bussiness "gomux_gorm/src/auth_module/bussiness/entities"
	repositories "gomux_gorm/src/auth_module/frameworks/repositories"
	token "gomux_gorm/src/auth_module/frameworks/token"
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

/*
* UserRepositorySpy
 */
type userRepositoryStructSpy struct{}

func (*userRepositoryStructSpy) FindByID(id int64) *tables.Users {
	return &tables.Users{}
}
func UserRepositorySpy() repositories.IUserRepository {
	return &userRepositoryStructSpy{}
}

/*
* PermissionRepositorySpy
 */
type permissionRepositoryStructSpy struct{}

func (*permissionRepositoryStructSpy) FindByID(id int64) *tables.Permissions {
	return &tables.Permissions{}
}
func PermissionRepositorySpy() repositories.IPermissionRepository {
	return &permissionRepositoryStructSpy{}
}

/*
* DecodedTokenSpy
 */
type decodedTokenStructSpy struct{}

func (*decodedTokenStructSpy) Decoded(t string) (*bussiness.TokenDecodedEntity, error) {
	return &bussiness.TokenDecodedEntity{}, nil
}
func DecodedTokenSpy() token.IDecodedToken {
	return &decodedTokenStructSpy{}
}

func TestAuthUsecase(t *testing.T) {
	userRepositorySpy := UserRepositorySpy()
	permissionRepositorySpy := PermissionRepositorySpy()
	decodedTokenSpy := DecodedTokenSpy()

	sut := AuthUsecase(&decodedTokenSpy, &userRepositorySpy, &permissionRepositorySpy)

	_, err := sut.Auth("", "")

	if err != nil {
		t.Errorf("Error")
	}

}
