package core

import "time"

type Permissions struct {
	ID          int64
	Role        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
