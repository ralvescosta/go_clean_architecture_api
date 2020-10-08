package sessioninterfaces

import (
	"encoding/json"
	"net/http"
	"strings"

	core "gomux_gorm/src/core/errors"
	usecases "gomux_gorm/src/m_session/application/usecases"
	bussiness "gomux_gorm/src/m_session/bussiness/entities"
)

// ISessionController ...
type ISessionController interface {
	Handle(res http.ResponseWriter, req *http.Request)
}

type controller struct {
	usecase *usecases.ISessionUsecase
}

// Handle ...
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

	host := strings.Split(req.Host, ":")
	session := bussiness.SessionEntity{
		Agent:         req.Header["User-Agent"][0],
		RemoteAddress: req.RemoteAddr,
		LocalAddress:  host[0],
		LocalPort:     host[1],
	}

	user, err := (*c.usecase).SessionUsecase(&body, &session)

	switch err.(type) {
	case *core.UnauthorizedError:
		res.WriteHeader(http.StatusUnauthorized)
		res.Write([]byte(`{"message": "User Credentials are wrong"}`))
		return
	case *core.ForbiddenError:
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte(`{"message": "User is not allowed"}`))
		return
	case *core.NotFoundError:
		res.WriteHeader(http.StatusNotFound)
		res.Write([]byte(`{"message": "User not found"}`))
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
	json.NewEncoder(res).Encode(user)
}

// SessionController ...
func SessionController(usecase *usecases.ISessionUsecase) ISessionController {
	return &controller{usecase}
}
