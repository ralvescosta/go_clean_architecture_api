package sessionbussinessentities

import (
	"time"
)

// UsersInput ...
type UsersInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SessionEntity ...
type SessionEntity struct {
	Name        string    `json:"name"`
	LastName    string    `json:"lastName"`
	Email       string    `json:"email"`
	AccessToken string    `json:"accessToken"`
	CreateAt    time.Time `json:"createAt"`
}
