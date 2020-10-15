package sessionframeworkscrypto

// import "golang.org/x/crypto/bcrypt"

/**/
// BcryptStruct ...
type BcryptStruct struct {
	CompareHashAndPassword func(hashedPassword []byte, password []byte) error
}

/**/

// IHasher ...
type IHasher interface {
	CheckPasswordHash(password, hash string) bool
}

type hash struct {
	bcrypt *BcryptStruct
}

// CheckPasswordHash ...
func (h *hash) CheckPasswordHash(password, hash string) bool {
	err := (*h.bcrypt).CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Hash ...
func Hash(bcrypt *BcryptStruct) IHasher {
	return &hash{bcrypt}
}
