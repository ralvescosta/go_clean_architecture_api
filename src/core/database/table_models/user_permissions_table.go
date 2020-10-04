package core

import (
	"time"

	"gorm.io/gorm"
)

// UsersPermissions ...
type UsersPermissions struct {
	ID int64 `gorm:"primaryKey"`

	UserID    int64
	UserName  string
	UserEmail string

	PermissionID   int64
	PermissionRole string

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Users       []Users       `gorm:"many2many:users;"`
	Permissions []Permissions `gorm:"many2many:permissions;"`
}
