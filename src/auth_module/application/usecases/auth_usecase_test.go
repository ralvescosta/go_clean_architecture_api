package authapplicationusecases

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	test "gomux_gorm/src/auth_module/application/__test__"
	bussiness "gomux_gorm/src/auth_module/bussiness/entities"
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
)

type mockedStruct struct {
	decodedTokenSpy         *test.DecodedTokenSpy
	userRepositorySpy       *test.UserRepositorySpy
	permissionRepositorySpy *test.PermissionRepositorySpy
	mockTokenDecoded        *bussiness.TokenDecodedEntity
	mockUser                *tables.Users
	mockPermission          *tables.Permissions
	mockAuthenticatedUser   *bussiness.AuthenticatedUser
}

func makeMocks() *mockedStruct {
	a := &test.DecodedTokenSpy{}
	b := &test.UserRepositorySpy{}
	c := &test.PermissionRepositorySpy{}

	mockUser := &tables.Users{
		ID:       1,
		Name:     "name",
		LastName: "last",
		Email:    "email",
		Password: "password",
	}

	mockPermission := &tables.Permissions{
		ID:   2,
		Role: "user",
	}

	mockTokenDecoded := &bussiness.TokenDecodedEntity{
		UserID:       mockUser.ID,
		PermissionID: mockPermission.ID,
	}

	mockAuthenticatedUser := &bussiness.AuthenticatedUser{
		ID:           1,
		Name:         mockUser.Name,
		Email:        mockUser.Email,
		PermissionID: mockPermission.ID,
	}

	return &mockedStruct{
		decodedTokenSpy:         a,
		userRepositorySpy:       b,
		permissionRepositorySpy: c,
		mockTokenDecoded:        mockTokenDecoded,
		mockUser:                mockUser,
		mockPermission:          mockPermission,
		mockAuthenticatedUser:   mockAuthenticatedUser,
	}
}

func TestAuthUsecase(t *testing.T) {
	mocks := makeMocks()

	mocks.decodedTokenSpy.On("Decoded", "some token").Return(mocks.mockTokenDecoded, nil)
	mocks.userRepositorySpy.On("FindByID", mocks.mockUser.ID).Return(mocks.mockUser)
	mocks.permissionRepositorySpy.On("FindByID", mocks.mockPermission.ID).Return(mocks.mockPermission)

	sut := AuthUsecase(mocks.decodedTokenSpy, mocks.userRepositorySpy, mocks.permissionRepositorySpy)

	result, _ := sut.Auth("some token", "user")

	assert.Equal(t, mocks.mockAuthenticatedUser, result, "Auth()")
}

func TestTokenExpiredError(t *testing.T) {
	mocks := makeMocks()

	mocks.decodedTokenSpy.On("Decoded", "some token").Return(&bussiness.TokenDecodedEntity{}, errors.New("Expired"))
	mocks.userRepositorySpy.On("FindByID", mocks.mockUser.ID).Return(mocks.mockUser)
	mocks.permissionRepositorySpy.On("FindByID", mocks.mockPermission.ID).Return(mocks.mockPermission)

	sut := AuthUsecase(mocks.decodedTokenSpy, mocks.userRepositorySpy, mocks.permissionRepositorySpy)

	_, err := sut.Auth("some token", "user")
	assert.Equal(t, "Something went wrong with the  request. Server returned 401 status.", err.Error(), "Should ReturnToken Expired Error If Token Decoded Return Error")
}

func TestPermissionNotAllowedError(t *testing.T) {
	mocks := makeMocks()

	mocks.decodedTokenSpy.On("Decoded", "some token").Return(mocks.mockTokenDecoded, nil)
	mocks.userRepositorySpy.On("FindByID", mocks.mockUser.ID).Return(mocks.mockUser)
	mocks.permissionRepositorySpy.On("FindByID", mocks.mockPermission.ID).Return(mocks.mockPermission)

	sut := AuthUsecase(mocks.decodedTokenSpy, mocks.userRepositorySpy, mocks.permissionRepositorySpy)

	_, err := sut.Auth("some token", "admin")
	assert.Equal(t, "Something went wrong with the  request. Server returned 401 status.", err.Error(), "Should Return Permission Not Allowed Error If User Role Not Allowed")
}
