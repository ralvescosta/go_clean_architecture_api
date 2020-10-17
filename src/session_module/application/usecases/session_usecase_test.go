package sessionapplicationusecases

import (
	"errors"
	"testing"

	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
	test "gomux_gorm/src/session_module/application/__test__"
	bussiness "gomux_gorm/src/session_module/bussiness/entities"

	"github.com/stretchr/testify/assert"
)

type mockedStruct struct {
	userRepositorySpy             *test.UserRepositorySpy
	sessionRepositorySpy          *test.SessionRepositorySpy
	usersPermissionsRepositorySpy *test.UsersPermissionsRepositorySpy
	hasherSpy                     *test.HasherSpy
	tokenSpy                      *test.CreateTokenSpy
	mockUser                      *tables.Users
	mockSessionEntity             *bussiness.SessionEntity
	mockPermission                *tables.Permissions
	mockUsersPermissions          *[]tables.UsersPermissions
	mockUsersInput                *bussiness.UsersInput
	mockUserSessionEntity         *bussiness.UserSessionEntity
}

func makeMocks() *mockedStruct {
	userRepositorySpy := &test.UserRepositorySpy{}
	sessionRepositorySpy := &test.SessionRepositorySpy{}
	usersPermissionsRepositorySpy := &test.UsersPermissionsRepositorySpy{}
	hasherSpy := &test.HasherSpy{}
	tokenSpy := &test.CreateTokenSpy{}

	mockUser := &tables.Users{
		ID:       1,
		Name:     "name",
		LastName: "last",
		Email:    "email",
		Password: "password",
	}

	mockSessionEntity := &bussiness.SessionEntity{
		Agent:         "agent",
		LocalAddress:  "local",
		LocalPort:     "port",
		RemoteAddress: "remote",
		AccessToken:   "some token",
	}

	mockPermission := &tables.Permissions{
		ID:   2,
		Role: "user",
	}

	mockUsersPermissions := &[]tables.UsersPermissions{
		tables.UsersPermissions{
			ID:             1,
			UserID:         mockUser.ID,
			UserName:       mockUser.Name,
			UserEmail:      mockUser.Email,
			PermissionID:   mockPermission.ID,
			PermissionRole: mockPermission.Role,
		},
	}

	mockUsersInput := &bussiness.UsersInput{
		Email:    mockUser.Email,
		Password: "123",
	}

	mockUserSessionEntity := &bussiness.UserSessionEntity{
		AccessToken: mockSessionEntity.AccessToken,
		Name:        mockUser.Name,
		LastName:    mockUser.LastName,
		Email:       mockUser.Email,
	}

	return &mockedStruct{
		userRepositorySpy:             userRepositorySpy,
		sessionRepositorySpy:          sessionRepositorySpy,
		usersPermissionsRepositorySpy: usersPermissionsRepositorySpy,
		hasherSpy:                     hasherSpy,
		tokenSpy:                      tokenSpy,
		mockUser:                      mockUser,
		mockSessionEntity:             mockSessionEntity,
		mockPermission:                mockPermission,
		mockUsersPermissions:          mockUsersPermissions,
		mockUsersInput:                mockUsersInput,
		mockUserSessionEntity:         mockUserSessionEntity,
	}
}

/**/

func TestSessionUsecase(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(mocks.mockUser)
	mocks.sessionRepositorySpy.On("Create", mocks.mockSessionEntity, mocks.mockUser)
	mocks.usersPermissionsRepositorySpy.On("FindUserPermissions", mocks.mockUser.ID).Return(mocks.mockUsersPermissions)
	mocks.hasherSpy.On("CheckPasswordHash", "123", mocks.mockUser.Password).Return(true)
	mocks.tokenSpy.On("CreateToken", &mocks.mockUser.ID, &mocks.mockPermission.ID).Return(mocks.mockSessionEntity.AccessToken, nil)

	sut := SessionUsecase(mocks.userRepositorySpy, mocks.sessionRepositorySpy, mocks.usersPermissionsRepositorySpy, mocks.hasherSpy, mocks.tokenSpy)

	result, _ := sut.SessionUsecase(mocks.mockUsersInput, mocks.mockSessionEntity)

	assert.Equal(t, mocks.mockUserSessionEntity, result, "SessionUsecase()")
}

func TestNotFoundError(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(&tables.Users{})
	mocks.sessionRepositorySpy.On("Create", mocks.mockSessionEntity, mocks.mockUser)
	mocks.usersPermissionsRepositorySpy.On("FindUserPermissions", mocks.mockUser.ID).Return(mocks.mockUsersPermissions)
	mocks.hasherSpy.On("CheckPasswordHash", "123", mocks.mockUser.Password).Return(true)
	mocks.tokenSpy.On("CreateToken", &mocks.mockUser.ID, &mocks.mockPermission.ID).Return(mocks.mockSessionEntity.AccessToken, nil)

	sut := SessionUsecase(mocks.userRepositorySpy, mocks.sessionRepositorySpy, mocks.usersPermissionsRepositorySpy, mocks.hasherSpy, mocks.tokenSpy)

	_, err := sut.SessionUsecase(mocks.mockUsersInput, mocks.mockSessionEntity)

	assert.Equal(t, "Something went wrong with the  request. Server returned 404 status.", err.Error(), "Should Return Not Found Error If User Email Not Registered")
}

func TestUnauthorizedError1(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(mocks.mockUser)
	mocks.sessionRepositorySpy.On("Create", mocks.mockSessionEntity, mocks.mockUser)
	mocks.usersPermissionsRepositorySpy.On("FindUserPermissions", mocks.mockUser.ID).Return(mocks.mockUsersPermissions)
	mocks.hasherSpy.On("CheckPasswordHash", "123", mocks.mockUser.Password).Return(false)
	mocks.tokenSpy.On("CreateToken", &mocks.mockUser.ID, &mocks.mockPermission.ID).Return(mocks.mockSessionEntity.AccessToken, nil)

	sut := SessionUsecase(mocks.userRepositorySpy, mocks.sessionRepositorySpy, mocks.usersPermissionsRepositorySpy, mocks.hasherSpy, mocks.tokenSpy)

	_, err := sut.SessionUsecase(mocks.mockUsersInput, mocks.mockSessionEntity)

	assert.Equal(t, "Something went wrong with the  request. Server returned 401 status.", err.Error(), "Should Return Unauthorized Error If Check Password Hash Returns False")
}

func TestUnauthorizedError2(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(mocks.mockUser)
	mocks.sessionRepositorySpy.On("Create", mocks.mockSessionEntity, mocks.mockUser)
	mocks.usersPermissionsRepositorySpy.On("FindUserPermissions", mocks.mockUser.ID).Return(&[]tables.UsersPermissions{tables.UsersPermissions{UserID: 1, PermissionID: 1, PermissionRole: "unauthorized"}})
	mocks.hasherSpy.On("CheckPasswordHash", "123", mocks.mockUser.Password).Return(true)
	mocks.tokenSpy.On("CreateToken", &mocks.mockUser.ID, &mocks.mockPermission.ID).Return(mocks.mockSessionEntity.AccessToken, nil)

	sut := SessionUsecase(mocks.userRepositorySpy, mocks.sessionRepositorySpy, mocks.usersPermissionsRepositorySpy, mocks.hasherSpy, mocks.tokenSpy)

	_, err := sut.SessionUsecase(mocks.mockUsersInput, mocks.mockSessionEntity)

	assert.Equal(t, "Something went wrong with the  request. Server returned 403 status.", err.Error(), "Should Return Unauthorized Error If Check Password Hash Returns False")
}

func TestInternalServerError(t *testing.T) {
	mocks := makeMocks()

	mocks.userRepositorySpy.On("FindByEmail", mocks.mockUser.Email).Return(mocks.mockUser)
	mocks.sessionRepositorySpy.On("Create", mocks.mockSessionEntity, mocks.mockUser)
	mocks.usersPermissionsRepositorySpy.On("FindUserPermissions", mocks.mockUser.ID).Return(mocks.mockUsersPermissions)
	mocks.hasherSpy.On("CheckPasswordHash", "123", mocks.mockUser.Password).Return(true)
	mocks.tokenSpy.On("CreateToken", &mocks.mockUser.ID, &mocks.mockPermission.ID).Return(mocks.mockSessionEntity.AccessToken, errors.New("Some Error"))

	sut := SessionUsecase(mocks.userRepositorySpy, mocks.sessionRepositorySpy, mocks.usersPermissionsRepositorySpy, mocks.hasherSpy, mocks.tokenSpy)

	_, err := sut.SessionUsecase(mocks.mockUsersInput, mocks.mockSessionEntity)

	assert.Equal(t, "Something went wrong with the  request. Server returned 500 status.", err.Error(), "Should Return Internal Server Error If Some Error Occur On CreateToken")
}
