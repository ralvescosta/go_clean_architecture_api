package signinframeworkscrypto

// import "golang.org/x/crypto/bcrypt"

/**/
// BcryptStruct ...
type BcryptStruct struct {
	GenerateFromPassword func(password []byte, cost int) ([]byte, error)
}

/**/

// IHasher ...
type IHasher interface {
	HashPassword(password string) (string, error)
}

type hash struct {
	bcrypt *BcryptStruct
}

// HashPassword ...
func (h *hash) HashPassword(password string) (string, error) {
	bytes, err := (*h.bcrypt).GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

// Hash ...
func Hash(bcrypt *BcryptStruct) IHasher {
	return &hash{bcrypt}
}
