package signinbussinessentities

// RegisterUsersEntity ...
type RegisterUsersEntity struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
