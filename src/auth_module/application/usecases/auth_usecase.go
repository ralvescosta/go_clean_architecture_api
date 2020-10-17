package authapplicationusecases

import (
	bussiness "gomux_gorm/src/auth_module/bussiness/entities"
	repositories "gomux_gorm/src/auth_module/frameworks/repositories"
	token "gomux_gorm/src/auth_module/frameworks/token"
	core "gomux_gorm/src/core_module/bussiness/errors"
)

type authUsecase struct {
	token                *token.IDecodedToken
	userRepository       *repositories.IUserRepository
	permissionRepository *repositories.IPermissionRepository
}

//Auth ..
func (u *authUsecase) Auth(token string, role string) (*bussiness.AuthenticatedUser, error) {

	decodedToken, err := (*u.token).Decoded(token)
	if err != nil {
		return nil, &core.TokenExpiredError{}
	}

	user := (*u.userRepository).FindByID(decodedToken.UserID)
	permission := (*u.permissionRepository).FindByID(decodedToken.PermissionID)

	if permission.Role != role {
		return nil, &core.PermissionNotAllowedError{}
	}

	return &bussiness.AuthenticatedUser{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PermissionID: decodedToken.PermissionID,
	}, nil
}

// AuthUsecase ...
func AuthUsecase(token token.IDecodedToken, userRepository repositories.IUserRepository, permissionRepository repositories.IPermissionRepository) IAuthUsecase {
	return &authUsecase{&token, &userRepository, &permissionRepository}
}
