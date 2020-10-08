package authframeworkstoken

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

// IToken ...
type IToken interface {
	Decoded(t string) (*TokenDecoded, error)
}

type token struct{}

// TokenDecoded ...
type TokenDecoded struct {
	userID       int64
	permissionID int64
}

// VerifyToken ...
func (*token) Decoded(t string) (*TokenDecoded, error) {

	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("ACCESS_SECRET"), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%i", claims["user_id"]), 10, 64)
		permissionID, err := strconv.ParseInt(fmt.Sprintf("%i", claims["permission_id"]), 10, 64)

		if err != nil {
			return nil, err
		}

		return &TokenDecoded{
			permissionID: permissionID,
			userID:       userID,
		}, nil

	}

	return nil, errors.New("Error")
}

// TokenConstructor ...
func TokenConstructor() IToken {
	return &token{}
}
