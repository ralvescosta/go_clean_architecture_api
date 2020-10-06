package sessionframeworkstoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// IToken ...
type IToken interface {
	CreateToken(userID int64) (string, error)
}

type token struct{}

// CreateToken ...
func (*token) CreateToken(userID int64) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return "", err
	}
	return token, nil
}

// TokenConstructor ...
func TokenConstructor() IToken {
	return &token{}
}
