package sessionapplicationusecases

import (
	"errors"

	bussiness "gomux_gorm/src/session_module/bussiness/entities"
	crypto "gomux_gorm/src/session_module/frameworks/crypto"
	repositories "gomux_gorm/src/session_module/frameworks/repositories"
	token "gomux_gorm/src/session_module/frameworks/token"
)

type usecase struct {
	repository *repositories.IUserRepository
	crypto     *crypto.IHasher
	token      *token.IToken
}

// ISessionUsecase ...
type ISessionUsecase interface {
	SessionUsecase(userInput *bussiness.UsersInput) (*bussiness.SessionEntity, error)
}

// SessionUsecase ...
func (u *usecase) SessionUsecase(userInput *bussiness.UsersInput) (*bussiness.SessionEntity, error) {

	user := (*u.repository).FindByEmail(userInput.Email)

	if user.ID == 0 {
		return nil, errors.New("user do not exist")
	}

	check := (*u.crypto).CheckPasswordHash(userInput.Password, user.Password)

	if !check {
		return nil, errors.New("Wrong Credentials")
	}

	token, err := (*u.token).CreateToken(user.ID)

	if err != nil {
		return nil, errors.New("JWT Error")
	}

	return &bussiness.SessionEntity{
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		AccessToken: token,
		CreateAt:    user.CreatedAt,
	}, nil
}

// SessionUsecaseConstructor ...
func SessionUsecaseConstructor(repository *repositories.IUserRepository, crypto *crypto.IHasher, token *token.IToken) ISessionUsecase {
	return &usecase{repository, crypto, token}
}
