package session_application_usecases

import (
	"errors"
	bussiness "gomux_gorm/src/session_module/bussiness/entities"
	crypto "gomux_gorm/src/session_module/frameworks/crypto"
	repositories "gomux_gorm/src/session_module/frameworks/repositories"
)

type usecase struct {
	repository *repositories.IUserRepository
	crypto     *crypto.IHasher
}

type ISessionUsecase interface {
	SessionUsecase(userInput *bussiness.UsersInput) (*bussiness.UsersEntity, error)
}

func (u *usecase) SessionUsecase(userInput *bussiness.UsersInput) (*bussiness.UsersEntity, error) {

	user := (*u.repository).FindByEmail(userInput.Email)

	if user.ID == 0 {
		return nil, errors.New("user do not exist")
	}

	check := (*u.crypto).CheckPasswordHash(userInput.Password, user.Password)

	if !check {
		return nil, errors.New("Wrong Credentials")
	}

	return &bussiness.UsersEntity{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		CreateAt: user.CreatedAt,
	}, nil
}

func SessionUsecaseConstructor(repository *repositories.IUserRepository, crypto *crypto.IHasher) ISessionUsecase {
	return &usecase{repository, crypto}
}
