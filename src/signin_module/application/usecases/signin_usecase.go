package signin_application_usecases

import (
	"errors"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"
	crypto "gomux_gorm/src/signin_module/frameworks/crypto"
	repositories "gomux_gorm/src/signin_module/frameworks/repositories"
)

type usecase struct {
	repository *repositories.IUserRepository
	crypto     *crypto.IHasher
}

type ISigninUsecase interface {
	SigninUsecase(user *bussiness.RegisterUsersEntity) error
}

func (u *usecase) SigninUsecase(user *bussiness.RegisterUsersEntity) error {

	userAlreadyRegistered := (*u.repository).FindByEmail(user.Email)

	if userAlreadyRegistered.ID != 0 {
		return errors.New("user already exist")
	}

	hashPassword, err := (*u.crypto).HashPassword(user.Password)

	if err != nil {
		return errors.New("Something Wrong in Hash password")
	}
	user.Password = hashPassword

	(*u.repository).Create(user)

	return nil
}

func SigninUsecaseConstructor(repository *repositories.IUserRepository, crypto *crypto.IHasher) ISigninUsecase {
	return &usecase{repository, crypto}
}
