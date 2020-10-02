package session_frameworks_crypto

import "golang.org/x/crypto/bcrypt"

type IHasher interface {
	CheckPasswordHash(password, hash string) bool
}

type hash struct{}

func (*hash) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashConstructor() IHasher {
	return &hash{}
}
