package authframeworkstoken

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

// IToken ...
type IToken interface {
	VerifyToken(t string) (*jwt.Token, error)
}

type token struct{}

// VerifyToken ...
func (*token) VerifyToken(t string) (*jwt.Token, error) {

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("ACCESS_SECRET"), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenConstructor ...
func TokenConstructor() IToken {
	return &token{}
}
