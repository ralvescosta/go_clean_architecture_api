package authframeworkstoken

import bussiness "gomux_gorm/src/auth_module/bussiness/entities"

// IDecodedToken ...
type IDecodedToken interface {
	Decoded(t string) (*bussiness.TokenDecodedEntity, error)
}
