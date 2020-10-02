package signin_interfaces

import (
	"encoding/json"
	usecases "gomux_gorm/src/signin_module/application/usecases"
	bussiness "gomux_gorm/src/signin_module/bussiness/entities"

	"net/http"
)

type ISigninController interface {
	Handle(res http.ResponseWriter, req *http.Request)
}

type controller struct {
	usecase *usecases.ISigninUsecase
}

func (c *controller) Handle(res http.ResponseWriter, req *http.Request) {

	var body bussiness.RegisterUsersEntity
	err := json.NewDecoder(req.Body).Decode(&body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "Internal Server Error"}`))
		return
	}

	if body.Name == "" || body.LastName == "" || body.Email == "" || body.Password == "" {
		res.WriteHeader(http.StatusUnsupportedMediaType)
		res.Write([]byte(`{"message": "Body wrong format"}`))
		return
	}

	err = (*c.usecase).SigninUsecase(&body)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "User Already Exist"}`))
		return
	}

	res.WriteHeader(http.StatusCreated)
	// json.NewEncoder(res).Encode("{}")
	res.Write([]byte(`{}`))
}

func SigninController(usecase *usecases.ISigninUsecase) ISigninController {
	return &controller{usecase}
}
