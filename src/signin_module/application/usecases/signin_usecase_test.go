package signinapplicationusecases

import (
	"errors"
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
	test "gomux_gorm/src/signin_module/application/__test__"
	entities "gomux_gorm/src/signin_module/bussiness/entities"

	"testing"

	"github.com/stretchr/testify/assert"
)

/*
* MAKE SUT
 */
type mockedStruct struct {
	userRepositorySpy            *test.UserRepositorySpy
	userPermissionsRepositorySpy *test.UserPermissionsRepositorySpy
	cryptoSpy                    *test.CryptoSpy
	mockUser                     *tables.Users
	mockPermission               *tables.Permissions
	mockRegisterUser             *entities.RegisterUsersEntity
	mockUsersPermissions         *tables.UsersPermissions
}

func makeMocks() *mockedStruct {
	userRepositorySpy := &test.UserRepositorySpy{}
	userPermissionsRepositorySpy := &test.UserPermissionsRepositorySpy{}
	cryptoSpy := &test.CryptoSpy{}

	mockUser := &tables.Users{
		ID:       1,
		Name:     "name",
		LastName: "last",
		Email:    "email",
		Password: "password",
	}

	mockPermission := &tables.Permissions{
		ID:   1,
		Role: "user",
	}

	mockRegisterUser := &entities.RegisterUsersEntity{
		Name:     mockUser.Name,
		LastName: mockUser.LastName,
		Email:    mockUser.Email,
		Password: "123",
	}

	mockUsersPermissions := &tables.UsersPermissions{
		ID:             1,
		UserID:         mockUser.ID,
		UserName:       mockUser.Name,
		UserEmail:      mockUser.Email,
		PermissionID:   mockPermission.ID,
		PermissionRole: mockPermission.Role,
	}

	return &mockedStruct{
		userRepositorySpy:            userRepositorySpy,
		userPermissionsRepositorySpy: userPermissionsRepositorySpy,
		cryptoSpy:                    cryptoSpy,
		mockUser:                     mockUser,
		mockPermission:               mockPermission,
		mockRegisterUser:             mockRegisterUser,
		mockUsersPermissions:         mockUsersPermissions,
	}
}

/**/

func TestSigninUsecase(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(&tables.Users{})
	mocks.cryptoSpy.On("HashPassword", mocks.mockRegisterUser.Password).Return("some hash", nil)
	mocks.userRepositorySpy.On("Create", mocks.mockRegisterUser).Return(mocks.mockUser)
	mocks.userPermissionsRepositorySpy.On("Create", mocks.mockUser, mocks.mockPermission.ID, mocks.mockPermission.Role).Return(mocks.mockUsersPermissions)

	sut := SigninUsecase(mocks.userRepositorySpy, mocks.userPermissionsRepositorySpy, mocks.cryptoSpy)

	err := sut.SigninUsecase(mocks.mockRegisterUser)

	assert.Equal(t, nil, err, "SigninUsecase()")
}

func TestConflictError(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(mocks.mockUser)
	mocks.cryptoSpy.On("HashPassword", mocks.mockRegisterUser.Password).Return("some hash", nil)
	mocks.userRepositorySpy.On("Create", mocks.mockRegisterUser).Return(mocks.mockUser)
	mocks.userPermissionsRepositorySpy.On("Create", mocks.mockUser, mocks.mockPermission.ID, mocks.mockPermission.Role).Return(mocks.mockUsersPermissions)

	sut := SigninUsecase(mocks.userRepositorySpy, mocks.userPermissionsRepositorySpy, mocks.cryptoSpy)

	err := sut.SigninUsecase(mocks.mockRegisterUser)

	assert.Equal(t, "Something went wrong with the  request. Server returned 409 status.", err.Error(), "Should Return Conflict Error If User Email Already Exist")
}

func TestInternalServerError(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(&tables.Users{})
	mocks.cryptoSpy.On("HashPassword", mocks.mockRegisterUser.Password).Return("some hash", errors.New("some error"))
	mocks.userRepositorySpy.On("Create", mocks.mockRegisterUser).Return(mocks.mockUser)
	mocks.userPermissionsRepositorySpy.On("Create", mocks.mockUser, mocks.mockPermission.ID, mocks.mockPermission.Role).Return(mocks.mockUsersPermissions)

	sut := SigninUsecase(mocks.userRepositorySpy, mocks.userPermissionsRepositorySpy, mocks.cryptoSpy)

	err := sut.SigninUsecase(mocks.mockRegisterUser)

	assert.Equal(t, "Something went wrong with the  request. Server returned 500 status.", err.Error(), "Should Return Internal Serber Error If Some Error Occur In Create HashPassword")
}
