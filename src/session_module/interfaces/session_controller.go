package session_interfaces

import (
	"encoding/json"
	usecases "gomux_gorm/src/session_module/application/usecases"
	bussiness "gomux_gorm/src/session_module/bussiness/entities"

	"net/http"
)

type ISessionController interface {
	Handle(res http.ResponseWriter, req *http.Request)
}

type controller struct {
	usecase *usecases.ISessionUsecase
}

func (c *controller) Handle(res http.ResponseWriter, req *http.Request) {

	var body bussiness.UsersInput
	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "Internal Server Error"}`))
		return
	}

	if body.Email == "" || body.Password == "" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		res.Write([]byte(`{"message": "Body wrong format"}`))
		return
	}

	user, err := (*c.usecase).SessionUsecase(&body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "User Credentials are wrong"}`))
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(user)
}

func SessionController(usecase *usecases.ISessionUsecase) ISessionController {
	return &controller{usecase}
}
