package sessionframeworkscrypto

import "golang.org/x/crypto/bcrypt"

// IHasher ...
type IHasher interface {
	CheckPasswordHash(password, hash string) bool
}

type hash struct{}

// CheckPasswordHash ...
func (*hash) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// HashConstructor ...
func HashConstructor() IHasher {
	return &hash{}
}
