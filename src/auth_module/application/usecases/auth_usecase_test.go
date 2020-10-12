package authapplicationusecases

import (
	"fmt"
	"testing"

	bussiness "gomux_gorm/src/auth_module/bussiness/entities"
	repositories "gomux_gorm/src/auth_module/frameworks/repositories"
	token "gomux_gorm/src/auth_module/frameworks/token"
	core "gomux_gorm/src/core_module/bussiness/errors"
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

/*
* UserRepositorySpy
 */
type userRepositoryStructSpy struct {
	res *tables.Users
}

func (*userRepositoryStructSpy) FindByID(id int64) *tables.Users {
	return &tables.Users{}
}
func UserRepositorySpy(userRepoRes *tables.Users) repositories.IUserRepository {
	return &userRepositoryStructSpy{res: userRepoRes}
}

/**/

/*
* PermissionRepositorySpy
 */
type permissionRepositoryStructSpy struct {
	res *tables.Permissions
}

func (p *permissionRepositoryStructSpy) FindByID(id int64) *tables.Permissions {
	return p.res
}
func PermissionRepositorySpy(permissionRepoRes *tables.Permissions) repositories.IPermissionRepository {
	return &permissionRepositoryStructSpy{res: permissionRepoRes}
}

/**/

/*
* DecodedTokenSpy
 */
type decodedTokenStructSpy struct {
	err error
}

func (d *decodedTokenStructSpy) Decoded(t string) (*bussiness.TokenDecodedEntity, error) {
	return &bussiness.TokenDecodedEntity{}, d.err
}
func DecodedTokenSpy(decodedTokenErr error) token.IDecodedToken {
	return &decodedTokenStructSpy{err: decodedTokenErr}
}

/**/

/*
* MAKE SUT
 */
func makeSut(userRepoRes *tables.Users, permissionRepoRes *tables.Permissions, decodedTokenErr error) IAuthUsecase {
	userRepositorySpy := UserRepositorySpy(userRepoRes)
	permissionRepositorySpy := PermissionRepositorySpy(permissionRepoRes)
	decodedTokenSpy := DecodedTokenSpy(decodedTokenErr)

	sut := AuthUsecase(&decodedTokenSpy, &userRepositorySpy, &permissionRepositorySpy)

	return sut
}

/**/

func TestAuthUsecase(t *testing.T) {
	sut := makeSut(&tables.Users{}, &tables.Permissions{}, nil)

	_, err := sut.Auth("", "")

	if err != nil {
		t.Errorf("AuthUsecase() wrong statement")
	}
}

func TestShouldReturnTokenExpiredErrorIfTokenDecodedRetornError(t *testing.T) {
	sut := makeSut(&tables.Users{}, &tables.Permissions{}, &core.TokenExpiredError{})

	_, err := sut.Auth("", "")
	fmt.Println(err)

	if err.Error() != "Something went wrong with the  request. Server returned 401 status." {
		t.Errorf("DecodedToken Unexpected Response")
	}
}

func TestShouldReturnPermissionNotAllowedErrorIfUserRoleNotAllowed(t *testing.T) {
	sut := makeSut(
		&tables.Users{},
		&tables.Permissions{
			ID:          1,
			Role:        "user",
			Description: "",
		},
		nil,
	)

	_, err := sut.Auth("", "admin")

	if err.Error() != "Something went wrong with the  request. Server returned 401 status." {
		t.Errorf("AuthUsecase role match Unexpected Response")
	}

}
