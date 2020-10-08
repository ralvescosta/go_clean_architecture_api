package sessionapplicationusecases

import (
	core "gomux_gorm/src/core/errors"
	bussiness "gomux_gorm/src/m_session/bussiness/entities"
	crypto "gomux_gorm/src/m_session/frameworks/crypto"
	repositories "gomux_gorm/src/m_session/frameworks/repositories"
	token "gomux_gorm/src/m_session/frameworks/token"
)

type usecase struct {
	userRepository             *repositories.IUserRepository
	sessionRepository          *repositories.ISessionRepository
	usersPermissionsRepository *repositories.IUsersPermissionsRepository
	crypto                     *crypto.IHasher
	token                      *token.ICreateToken
}

// ISessionUsecase ...
type ISessionUsecase interface {
	SessionUsecase(userInput *bussiness.UsersInput, session *bussiness.SessionEntity) (*bussiness.UserSessionEntity, error)
}

// SessionUsecase ...
func (u *usecase) SessionUsecase(userInput *bussiness.UsersInput, session *bussiness.SessionEntity) (*bussiness.UserSessionEntity, error) {
	user := (*u.userRepository).FindByEmail(userInput.Email)
	if user.ID == 0 {
		return nil, &core.NotFoundError{}
	}

	check := (*u.crypto).CheckPasswordHash(userInput.Password, user.Password)
	if !check {
		return nil, &core.UnauthorizedError{}
	}

	userPermissions := (*u.usersPermissionsRepository).FindUserPermissions(user.ID)

	var userPermissionID *int64

	var err error = nil
	for _, element := range *userPermissions {
		if element.PermissionID == 0 || element.PermissionID == 1 {
			err = &core.ForbiddenError{}
			break
		}
		userPermissionID = &element.PermissionID
	}

	if err != nil {
		return nil, err
	}

	token, err := (*u.token).CreateToken(&user.ID, userPermissionID)
	if err != nil {
		return nil, &core.InternalServerError{}
	}
	session.AccessToken = token

	(*u.sessionRepository).Create(session, user)

	return &bussiness.UserSessionEntity{
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		AccessToken: token,
		CreateAt:    user.CreatedAt,
	}, nil
}

// SessionUsecase ...
func SessionUsecase(
	userRepository *repositories.IUserRepository,
	sessionRepository *repositories.ISessionRepository,
	usersPermissionsRepository *repositories.IUsersPermissionsRepository,
	crypto *crypto.IHasher,
	token *token.ICreateToken,
) ISessionUsecase {
	return &usecase{userRepository, sessionRepository, usersPermissionsRepository, crypto, token}
}
