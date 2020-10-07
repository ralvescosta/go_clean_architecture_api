package authinterfaces

import (
	"fmt"
	"net/http"
	"strings"
)

// Handle ...
func Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearToken := r.Header.Get("Authorization")
		strArr := strings.Split(bearToken, " ")
		if len(strArr) == 2 {
			fmt.Println(strArr[1])
			return
		}
	})
}
