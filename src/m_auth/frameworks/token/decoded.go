package authframeworkstoken

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/dgrijalva/jwt-go"

	bussiness "gomux_gorm/src/m_auth/bussiness/entities"
)

// IDecodedToken ...
type IDecodedToken interface {
	Decoded(t string) (*bussiness.TokenDecodedEntity, error)
}

type decodedToken struct{}

// VerifyToken ...
func (*decodedToken) Decoded(t string) (*bussiness.TokenDecodedEntity, error) {

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
		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		permissionID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["permission_id"]), 10, 64)

		if err != nil {
			return nil, err
		}

		return &bussiness.TokenDecodedEntity{
			UserID:       userID,
			PermissionID: permissionID,
		}, nil

	}

	return nil, errors.New("Some Error occur when Decoded Token")
}

// DecodedToken ...
func DecodedToken() IDecodedToken {
	return &decodedToken{}
}
