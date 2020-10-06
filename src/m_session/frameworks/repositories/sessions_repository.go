package sessionframeworksrepositories

import (
	"github.com/jinzhu/gorm"

	tables "gomux_gorm/src/core/database/table_models"
	bussiness "gomux_gorm/src/m_session/bussiness/entities"
)

type sessionRepository struct {
	db *gorm.DB
}

// ISessionRepository ...
type ISessionRepository interface {
	Create(session *bussiness.SessionEntity, user *tables.Users)
}

func (r *sessionRepository) Create(session *bussiness.SessionEntity, user *tables.Users) {
	r.db.Create(&tables.Sessions{
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
