package core

import "time"

// UsersPermissions ...
type UsersPermissions struct {
	ID int64

	UserID    int64
	UserName  string
	UserEmail string

	PermissionID   int64
	PermissionRole string

	CreatedAt time.Time
	UpdatedAt time.Time
}
