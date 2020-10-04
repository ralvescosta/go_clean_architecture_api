package core

import "time"

// Sessions ...
type Sessions struct {
	ID int64

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
}
