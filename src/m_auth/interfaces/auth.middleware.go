package authinterfaces

import (
	"fmt"
	"net/http"
	"strings"

	usecases "gomux_gorm/src/m_auth/application/usecases"
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

		fmt.Println(result, err)
		/*
		* JWT Framework
		 */

	})
}

// AuthMiddleware ...
func AuthMiddleware(usecase *usecases.IAuthUsecase) IAuthMiddleware {
	return &middleware{usecase}
}
