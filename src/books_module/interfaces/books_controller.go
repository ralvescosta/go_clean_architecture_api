package booksinterfaces

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
)

// IBooksController ...
type IBooksController interface {
	Handle(res http.ResponseWriter, req *http.Request)
}

type controller struct{}

func (*controller) Handle(res http.ResponseWriter, req *http.Request) {
	auth := context.Get(req, "auth")
	fmt.Println(auth)
	fmt.Fprintf(res, "Hello World")
}

// BooksController ...
func BooksController() IBooksController {
	return &controller{}
}
