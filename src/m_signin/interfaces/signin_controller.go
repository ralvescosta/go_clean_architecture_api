package signininterfaces

import (
	"encoding/json"
	"net/http"

	core "gomux_gorm/src/core/errors"
	usecases "gomux_gorm/src/m_signin/application/usecases"
	bussiness "gomux_gorm/src/m_signin/bussiness/entities"
)

// ISigninController ...
type ISigninController interface {
	Handle(res http.ResponseWriter, req *http.Request)
}

type controller struct {
	usecase *usecases.ISigninUsecase
}

// Handle ...
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

	switch err.(type) {
	case *core.UnauthorizedError:
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte(`{"message": "User Credentials are wrong"}`))
		return
	case *core.ForbiddenError:
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte(`{"message": "User is not allowed"}`))
		return
	case *core.ConflictError:
		res.WriteHeader(http.StatusConflict)
		res.Write([]byte(`{"message": "User Already exist"}`))
		return
	case *core.UnsupportedMediaTypeError:
		res.WriteHeader(http.StatusUnsupportedMediaType)
		res.Write([]byte(`{"message": "Wrong Body format"}`))
		return
	case *core.InternalServerError:
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "Internal Server Error"}`))
		return
	}

	res.WriteHeader(http.StatusCreated)
	res.Write([]byte(`{}`))
}

// SigninController ...
func SigninController(usecase *usecases.ISigninUsecase) ISigninController {
	return &controller{usecase}
}
