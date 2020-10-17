package sessionapplicationusecases

import bussiness "gomux_gorm/src/session_module/bussiness/entities"

// ISessionUsecase ...
type ISessionUsecase interface {
	SessionUsecase(userInput *bussiness.UsersInput, session *bussiness.SessionEntity) (*bussiness.UserSessionEntity, error)
}
