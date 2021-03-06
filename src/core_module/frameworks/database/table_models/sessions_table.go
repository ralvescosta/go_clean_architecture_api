package coreframeworksdatabase

import (
	"time"

	"gorm.io/gorm"
)

// Sessions ...
type Sessions struct {
	ID int64 `gorm:"primaryKey"`

	UserID    int64
	UserName  string
	UserEmail string

	PermissionID   int64
	PermissionRole string

	Agent         string
	RemoteAddress string
	LocalAddress  string
	LocalPort     string
	AccessToken   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
