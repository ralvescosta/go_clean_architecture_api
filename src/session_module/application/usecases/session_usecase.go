package sessionapplicationusecases

import (
	"errors"

	bussiness "gomux_gorm/src/session_module/bussiness/entities"
	crypto "gomux_gorm/src/session_module/frameworks/crypto"
	repositories "gomux_gorm/src/session_module/frameworks/repositories"
	token "gomux_gorm/src/session_module/frameworks/token"
)

type usecase struct {
	userRepository    *repositories.IUserRepository
	sessionRepository *repositories.ISessionRepository
	crypto            *crypto.IHasher
	token             *token.IToken
}

// ISessionUsecase ...
type ISessionUsecase interface {
	SessionUsecase(userInput *bussiness.UsersInput, session *bussiness.SessionEntity) (*bussiness.UserSessionEntity, error)
}

// SessionUsecase ...
func (u *usecase) SessionUsecase(userInput *bussiness.UsersInput, session *bussiness.SessionEntity) (*bussiness.UserSessionEntity, error) {
	user := (*u.userRepository).FindByEmail(userInput.Email)

	if user.ID == 0 {
		return nil, errors.New("user do not exist")
	}

	check := (*u.crypto).CheckPasswordHash(userInput.Password, user.Password)

	if !check {
		return nil, errors.New("Wrong Credentials")
	}

	token, err := (*u.token).CreateToken(user.ID)
	session.AccessToken = token

	(*u.sessionRepository).Create(session, user)

	if err != nil {
		return nil, errors.New("JWT Error")
	}

	return &bussiness.UserSessionEntity{
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		AccessToken: token,
		CreateAt:    user.CreatedAt,
	}, nil
}

// SessionUsecaseConstructor ...
func SessionUsecaseConstructor(
	userRepository *repositories.IUserRepository,
	sessionRepository *repositories.ISessionRepository,
	crypto *crypto.IHasher, token *token.IToken,
) ISessionUsecase {
	return &usecase{userRepository, sessionRepository, crypto, token}
}
