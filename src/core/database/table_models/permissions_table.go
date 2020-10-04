package core

import (
	"time"

	"gorm.io/gorm"
)

// Permissions ...
type Permissions struct {
	ID          int64
	Role        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
