package authapplicationusecases

import bussiness "gomux_gorm/src/auth_module/bussiness/entities"

// IAuthUsecase ...
type IAuthUsecase interface {
	Auth(token string, role string) (*bussiness.AuthenticatedUser, error)
}
