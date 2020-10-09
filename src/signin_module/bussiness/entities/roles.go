package signinbussinessentities

// Roles ...
const (
	RoleUnauthorized = iota
	RoleUser         = iota
	RoleAdmin        = iota
)

// Permissions ...
var Permissions = [3]string{"unauthorized ", "user", "admin"}
