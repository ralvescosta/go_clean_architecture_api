package signinapplicationusecases

import (
	core "gomux_gorm/src/core_module/bussiness/errors"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"
	crypto "gomux_gorm/src/signin_module/frameworks/crypto"
	repositories "gomux_gorm/src/signin_module/frameworks/repositories"
)

type usecase struct {
	userRepository             *repositories.IUserRepository
	usersPermissionsRepository *repositories.IUsersPermissionsRepository
	crypto                     *crypto.IHasher
}

// SigninUsecase ......
func (u *usecase) SigninUsecase(user *bussiness.RegisterUsersEntity) error {

	userAlreadyRegistered := (*u.userRepository).FindByEmail(user.Email)

	if userAlreadyRegistered.ID != 0 {
		return &core.ConflictError{}
	}

	hashPassword, err := (*u.crypto).HashPassword(user.Password)

	if err != nil {
		return &core.InternalServerError{}
	}
	user.Password = hashPassword

	createdUser := (*u.userRepository).Create(user)
	(*u.usersPermissionsRepository).Create(createdUser, bussiness.RoleUser, bussiness.Permissions[bussiness.RoleUser])

	return nil
}

// SigninUsecase ...
func SigninUsecase(userRepository repositories.IUserRepository, usersPermissionsRepository repositories.IUsersPermissionsRepository, crypto crypto.IHasher) ISigninUsecase {
	return &usecase{&userRepository, &usersPermissionsRepository, &crypto}
}
