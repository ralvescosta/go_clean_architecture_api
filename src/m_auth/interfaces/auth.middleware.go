package authinterfaces

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// Handle ...
func Handle(next http.Handler) http.Handler {
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

		/*
		* JWT Framework
		 */
		token, err := jwt.Parse(bearer, func(token *jwt.Token) (interface{}, error) {
			//Make sure that the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("ACCESS_SECRET"), nil
		})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(token)

		claims, ok := token.Claims.(jwt.MapClaims)

		if ok && token.Valid {

			userID, ok := claims["user_id"]
			permissionID, ok := claims["permission_id"]

			if !ok {
				fmt.Println(err)
			}

			fmt.Println(userID, permissionID)

		}

	})
}
