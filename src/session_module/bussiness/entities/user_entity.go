package session_bussiness_entities

import (
	"time"
)

type UsersInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersEntity struct {
	Name     string    `json:"name"`
	LastName string    `json:"lastName"`
	Email    string    `json:"email"`
	CreateAt time.Time `json:"createAt"`
}
