package signin_frameworks_crypto

import "golang.org/x/crypto/bcrypt"

type IHasher interface {
	HashPassword(password string) (string, error)
}

type hash struct{}

func (*hash) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func HashConstructor() IHasher {
	return &hash{}
}

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }
