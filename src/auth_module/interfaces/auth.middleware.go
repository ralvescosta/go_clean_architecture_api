package authinterfaces

import (
	"net/http"
	"strings"

	"github.com/gorilla/context"

	usecases "gomux_gorm/src/auth_module/application/usecases"
)

// IAuthMiddleware ...
type IAuthMiddleware interface {
	Handle(next http.Handler) http.Handler
}

type middleware struct {
	usecase *usecases.IAuthUsecase
}

// Handle ...
func (m *middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		bearerToken := req.Header.Get("Authorization")
		if bearerToken == "" {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte(`{"message": "User Credentials are wrong"}`))
			return
		}

		strArr := strings.Split(bearerToken, " ")
		if len(strArr) < 2 {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte(`{"message": "User Credentials are wrong"}`))
			return
		}

		bearer := strArr[1]
		result, err := (*m.usecase).Auth(bearer)
		if err != nil {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte(`{"message": "User Token expired"}`))
			return
		}

		context.Set(req, "auth", result)

		next.ServeHTTP(res, req)
	})
}

// AuthMiddleware ...
func AuthMiddleware(usecase *usecases.IAuthUsecase) IAuthMiddleware {
	return &middleware{usecase}
}
