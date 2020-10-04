package sessionframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	migrations "gomux_gorm/src/core/database/table_models"
	bussiness "gomux_gorm/src/session_module/bussiness/entities"
)

type sessionRepository struct {
	db *gorm.DB
}

// ISessionRepository ...
type ISessionRepository interface {
	Create(session *bussiness.SessionEntity, user *migrations.Users)
}

func (r *sessionRepository) Create(session *bussiness.SessionEntity, user *migrations.Users) {
	r.db.Create(&migrations.Sessions{
		UserID:         user.ID,
		UserName:       user.Name,
		UserEmail:      user.Email,
		PermissionID:   1,
		PermissionRole: "Role",
		Agent:          session.Agent,
		RemoteAddress:  session.RemoteAddress,
		LocalAddress:   session.LocalAddress,
		LocalPort:      session.LocalPort,
		AccessToken:    session.AccessToken,
	})
}

// SessionRepositoryConstructor ...
func SessionRepositoryConstructor(db *gorm.DB) ISessionRepository {
	return &sessionRepository{db}
}
