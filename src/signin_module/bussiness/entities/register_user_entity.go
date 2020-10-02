package signin_bussiness_entities

type RegisterUsersEntity struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
