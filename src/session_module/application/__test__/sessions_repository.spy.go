package test

import (
	"fmt"
	tables "gomux_gorm/src/core_module/frameworks/database/table_models"
	bussiness "gomux_gorm/src/session_module/bussiness/entities"

	"github.com/stretchr/testify/mock"
)

// SessionRepositorySpy ...
type SessionRepositorySpy struct {
	mock.Mock
}

// Create ...
func (s *SessionRepositorySpy) Create(session *bussiness.SessionEntity, user *tables.Users) {
	args := s.Called(session, user)
	fmt.Println(args)
}
