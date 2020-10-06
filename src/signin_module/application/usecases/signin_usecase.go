package signinapplicationusecases

import (
	core "gomux_gorm/src/core/errors"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"
	crypto "gomux_gorm/src/signin_module/frameworks/crypto"
	repositories "gomux_gorm/src/signin_module/frameworks/repositories"
)

type usecase struct {
	userRepository             *repositories.IUserRepository
	usersPermissionsRepository *repositories.IUsersPermissionsRepository
	crypto                     *crypto.IHasher
}

// ISigninUsecase ...
type ISigninUsecase interface {
	SigninUsecase(user *bussiness.RegisterUsersEntity) error
}

// SigninUsecase ...
func (u *usecase) SigninUsecase(user *bussiness.RegisterUsersEntity) error {

	userAlreadyRegistered := (*u.userRepository).FindByEmail(user.Email)

	if userAlreadyRegistered.ID != 0 {
		return &core.ConflictError{}
	}

	hashPassword, err := (*u.crypto).HashPassword(user.Password)

	if err != nil {
		return &core.UnauthorizedError{}
	}
	user.Password = hashPassword

	createdUser := (*u.userRepository).Create(user)
	(*u.usersPermissionsRepository).Create(createdUser, bussiness.RoleUser, bussiness.Permissions[bussiness.RoleUser])

	return nil
}

// SigninUsecaseConstructor ...
func SigninUsecaseConstructor(userRepository *repositories.IUserRepository, usersPermissionsRepository *repositories.IUsersPermissionsRepository, crypto *crypto.IHasher) ISigninUsecase {
	return &usecase{userRepository, usersPermissionsRepository, crypto}
}
