package authapplicationusecases

import (
	bussiness "gomux_gorm/src/auth_module/bussiness/entities"
	repositories "gomux_gorm/src/auth_module/frameworks/repositories"
	token "gomux_gorm/src/auth_module/frameworks/token"
)

type authUsecase struct {
	token          *token.IDecodedToken
	userRepository *repositories.IUserRepository
}

// IAuthUsecase ...
type IAuthUsecase interface {
	Auth(token string) (*bussiness.AuthenticatedUser, error)
}

//Auth ..
func (u *authUsecase) Auth(token string) (*bussiness.AuthenticatedUser, error) {

	decodedToken, err := (*u.token).Decoded(token)
	if err != nil {
		return nil, err
	}

	user := (*u.userRepository).FindByID(decodedToken.UserID)

	return &bussiness.AuthenticatedUser{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PermissionID: decodedToken.PermissionID,
	}, nil
}

// AuthUsecase ...
func AuthUsecase(token *token.IDecodedToken, userRepository *repositories.IUserRepository) IAuthUsecase {
	return &authUsecase{token, userRepository}
}
