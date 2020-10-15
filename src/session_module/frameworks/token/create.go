package sessionframeworkstoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

/**/

// JwtStruct ...
type JwtStruct struct {
	MapClaims          jwt.MapClaims
	NewWithClaims      func(method jwt.SigningMethod, claims jwt.Claims) *jwt.Token
	SigningMethodHS256 *jwt.SigningMethodHMAC
}

/**/

// ICreateToken ...
type ICreateToken interface {
	CreateToken(userID *int64, permissionID *int64) (string, error)
}

type createToken struct {
	jwt *JwtStruct
}

// CreateToken ...
func (j *createToken) CreateToken(userID *int64, permissionID *int64) (string, error) {
	var err error

	atClaims := (*j.jwt).MapClaims
	atClaims["authorized"] = true
	atClaims["user_id"] = userID
	atClaims["permission_id"] = permissionID
	atClaims["exp"] = time.Now().Add(time.Hour * 8).Unix()
	at := (*j.jwt).NewWithClaims((*j.jwt).SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte("ACCESS_SECRET"))
	if err != nil {
		return "", err
	}
	return token, nil
}

// CreateToken ...
func CreateToken(jwt *JwtStruct) ICreateToken {
	return &createToken{jwt}
}
