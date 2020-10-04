package sessioninterfaces

import (
	"encoding/json"
	"net/http"
	"strings"

	usecases "gomux_gorm/src/session_module/application/usecases"
	bussiness "gomux_gorm/src/session_module/bussiness/entities"
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

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"message": "User Credentials are wrong"}`))
		return
	}

	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(user)
}

// SessionController ...
func SessionController(usecase *usecases.ISessionUsecase) ISessionController {
	return &controller{usecase}
}
