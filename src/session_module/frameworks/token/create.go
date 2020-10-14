package sessionframeworkstoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

/**/

/**/

// ICreateToken ...
type ICreateToken interface {
	CreateToken(userID *int64, permissionID *int64) (string, error)
}

type createToken struct{}

// CreateToken ...
func (j *createToken) CreateToken(userID *int64, permissionID *int64) (string, error) {
	var err error

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["permission_id"] = permissionID
	atClaims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return "", err
	}
	return token, nil
}

// CreateToken ...
func CreateToken() ICreateToken {
	return &createToken{}
}
