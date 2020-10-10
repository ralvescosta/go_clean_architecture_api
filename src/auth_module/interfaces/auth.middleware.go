package authinterfaces

import (
	"net/http"
	"strings"

	"github.com/gorilla/context"

	usecases "gomux_gorm/src/auth_module/application/usecases"
	core "gomux_gorm/src/core_module/bussiness/errors"
)

// IAuthMiddleware ...
type IAuthMiddleware interface {
	Auth(role string) func(next http.Handler) http.Handler
}

type middleware struct {
	usecase *usecases.IAuthUsecase
}

// Handle ...
func (m *middleware) Auth(role string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			bearerToken := req.Header.Get("Authorization")
			if bearerToken == "" {
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"message": "Token not provide"}`))
				return
			}

			strArr := strings.Split(bearerToken, " ")
			if len(strArr) < 2 {
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"message": "Wrong Token format"}`))
				return
			}

			bearer := strArr[1]
			result, err := (*m.usecase).Auth(bearer, role)
			switch err.(type) {
			case *core.TokenExpiredError:
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"message": "Token Expired"}`))
				return
			case *core.PermissionNotAllowedError:
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"message": "Permission Not Allowed"}`))
				return
			}

			context.Set(req, "auth", result)

			next.ServeHTTP(res, req)
		})
	}
}

// AuthMiddleware ...
func AuthMiddleware(usecase *usecases.IAuthUsecase) IAuthMiddleware {
	return &middleware{usecase}
}
