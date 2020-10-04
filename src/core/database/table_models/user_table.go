package core

import (
	"time"

	"gorm.io/gorm"
)

// Users ...
type Users struct {
	ID        int64
	Name      string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
