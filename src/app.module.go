package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	coreDatabase "gomux_gorm/src/core_module/frameworks/database"
	coreMiddleware "gomux_gorm/src/core_module/interfaces"

	signinUsecases "gomux_gorm/src/signin_module/application/usecases"
	signinCrypto "gomux_gorm/src/signin_module/frameworks/crypto"
	signinRepositories "gomux_gorm/src/signin_module/frameworks/repositories"
	signinControllers "gomux_gorm/src/signin_module/interfaces"

	sessionUsecases "gomux_gorm/src/session_module/application/usecases"
	sessionCrypto "gomux_gorm/src/session_module/frameworks/crypto"
	sessionRepositories "gomux_gorm/src/session_module/frameworks/repositories"
	sessionToken "gomux_gorm/src/session_module/frameworks/token"
	sessionControllers "gomux_gorm/src/session_module/interfaces"

	authUsecases "gomux_gorm/src/auth_module/application/usecases"
	authRepositories "gomux_gorm/src/auth_module/frameworks/repositories"
	authToken "gomux_gorm/src/auth_module/frameworks/token"
	authMiddleware "gomux_gorm/src/auth_module/interfaces"

	booksControllers "gomux_gorm/src/books_module/interfaces"
)

type module struct {
	conn *gorm.DB
}

// IHttpServer ...
type IHttpServer interface {
	StartHTTPServer()
}

// StartHTTPServer ...
func (m *module) StartHTTPServer() {
	const PORT string = ":4000"

	conn := coreDatabase.ConnectToDatabase()
	defer conn.Close()

	m.conn = conn

	router := mux.NewRouter()

	router.Use(coreMiddleware.HeadersMiddleware)

	m.registerRouters(router)

	srv := &http.Server{
		Handler:      handlers.CompressHandler(router),
		Addr:         "127.0.0.1" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server Listening on port:", PORT)
	log.Fatalln(srv.ListenAndServe())
}

func (m *module) registerRouters(router *mux.Router) {

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "{\"status\": \"ok\"}")
	})

	_signinCrypto := signinCrypto.Hash()
	_signinUserRepository := signinRepositories.UserRepository(m.conn)
	_signinUsersPermissionsRepository := signinRepositories.UsersPermissionsRepository(m.conn)
	_signinUsecase := signinUsecases.SigninUsecase(&_signinUserRepository, &_signinUsersPermissionsRepository, &_signinCrypto)
	_signinController := signinControllers.SigninController(&_signinUsecase)
	router.HandleFunc("/signin", _signinController.Handle).Methods("POST")

	_sessionCrypto := sessionCrypto.Hash()
	_sessionCreateToken := sessionToken.CreateToken()
	_sessionUserRepository := sessionRepositories.UserRepository(m.conn)
	_sessionSessionRepository := sessionRepositories.SessionRepository(m.conn)
	_sessionUsersPermissionsRepository := sessionRepositories.UsersPermissionsRepository(m.conn)
	_sessionUsecase := sessionUsecases.SessionUsecase(&_sessionUserRepository, &_sessionSessionRepository, &_sessionUsersPermissionsRepository, &_sessionCrypto, &_sessionCreateToken)
	_sessionController := sessionControllers.SessionController(&_sessionUsecase)
	router.HandleFunc("/session", _sessionController.Handle).Methods("POST")

	_authToken := authToken.DecodedToken()
	_authUserRepository := authRepositories.UserRepository(m.conn)
	_authPermissionRepository := authRepositories.PermissionRepository(m.conn)
	_authUsecase := authUsecases.AuthUsecase(&_authToken, &_authUserRepository, &_authPermissionRepository)
	_authMiddleware := authMiddleware.AuthMiddleware(&_authUsecase)

	_booksController := booksControllers.BooksController()
	booksGroup := router.PathPrefix("/books").Subrouter()
	booksGroup.Use(_authMiddleware.Auth("admin"))
	booksGroup.HandleFunc("", _booksController.Handle).Methods("GET")

}

// HTTPServerController ...
func HTTPServerController() IHttpServer {
	return &module{}
}
