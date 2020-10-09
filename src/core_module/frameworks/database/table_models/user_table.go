package coreframeworksdatabase

import (
	"time"

	"gorm.io/gorm"
)

// Users ...
type Users struct {
	ID        int64 `gorm:"primaryKey"`
	Name      string
	LastName  string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
