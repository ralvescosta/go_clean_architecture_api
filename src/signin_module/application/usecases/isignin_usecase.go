package signinapplicationusecases

import bussiness "gomux_gorm/src/signin_module/bussiness/entities"

// ISigninUsecase ...
type ISigninUsecase interface {
	SigninUsecase(user *bussiness.RegisterUsersEntity) error
}
