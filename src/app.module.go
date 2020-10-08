package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	core "gomux_gorm/src/core/database"

	signinUsecases "gomux_gorm/src/m_signin/application/usecases"
	signinCrypto "gomux_gorm/src/m_signin/frameworks/crypto"
	signinRepositories "gomux_gorm/src/m_signin/frameworks/repositories"
	signinControllers "gomux_gorm/src/m_signin/interfaces"

	sessionUsecases "gomux_gorm/src/m_session/application/usecases"
	sessionCrypto "gomux_gorm/src/m_session/frameworks/crypto"
	sessionRepositories "gomux_gorm/src/m_session/frameworks/repositories"
	sessionToken "gomux_gorm/src/m_session/frameworks/token"
	sessionControllers "gomux_gorm/src/m_session/interfaces"

	authMiddleware "gomux_gorm/src/m_auth/interfaces"
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

	conn := core.ConnectToDatabase()
	defer conn.Close()

	m.conn = conn

	router := mux.NewRouter()

	router.Use(headersMiddleware)

	m.registerRouters(router)

	srv := &http.Server{
		Handler:      handlers.CompressHandler(router),
		Addr:         "127.0.0.1" + PORT,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Server Listening on port: ", PORT)
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
	_sessionToken := sessionToken.Token()
	_sessionUserRepository := sessionRepositories.UserRepository(m.conn)
	_sessionSessionRepository := sessionRepositories.SessionRepository(m.conn)
	_sessionUsersPermissionsRepository := sessionRepositories.UsersPermissionsRepository(m.conn)
	_sessionUsecase := sessionUsecases.SessionUsecase(&_sessionUserRepository, &_sessionSessionRepository, &_sessionUsersPermissionsRepository, &_sessionCrypto, &_sessionToken)
	_sessionController := sessionControllers.SessionController(&_sessionUsecase)
	router.HandleFunc("/session", _sessionController.Handle).Methods("POST")

	booksGroup := router.PathPrefix("/books").Subrouter()
	booksGroup.Use(authMiddleware.Handle)
	booksGroup.HandleFunc("", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World")
	}).Methods("GET")
	router.HandleFunc("/books2", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "Hello World - 2")
	}).Methods("GET")

}

func headersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		w.Header().Add("X-DNS-Prefetch-Control", "off")
		w.Header().Add("X-Frame-Options", "SAMEORIGIN")
		w.Header().Add("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
		w.Header().Add("X-Download-Options", "noopen")
		w.Header().Add("X-Content-Type-Options", "nosniff")
		w.Header().Add("X-XSS-Protection", "1; mode=block")
		w.Header().Add("Content-Security-Policy", "default-src 'none'")
		w.Header().Add("X-Content-Security-Policy", "default-src 'none'")
		w.Header().Add("X-WebKit-CSP", "default-src 'none'")
		w.Header().Add("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Add("Referrer-Policy", "origin-when-cross-origin,strict-origin-when-cross-origin")
		w.Header().Add("Access-Control-Allow-Origin	", "*")
		w.Header().Add("Vary", "Accept-Encoding")

		next.ServeHTTP(w, r)
	})
}

// HTTPServerController ...
func HTTPServerController() IHttpServer {
	return &module{}
}
