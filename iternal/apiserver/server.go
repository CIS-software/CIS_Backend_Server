package apiserver

import (
	"CIS_Backend_Server/iternal/handlers/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	router    *mux.Router
	log       *logrus.Logger
	handler   *handlers.Handlers
	secretKey string
}

func serverNew(log *logrus.Logger, router *mux.Router, handler *handlers.Handlers, secretKey string) *Server {
	s := &Server{
		router:    router,
		log:       log,
		handler:   handler,
		secretKey: secretKey,
	}

	s.configureRouter()

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {
	//users URLs
	s.router.HandleFunc("/user/{id}", s.handler.Users().Get()).Methods("GET")
	s.router.HandleFunc("/create-user", s.handler.Users().Create()).Methods("POST")
	s.router.HandleFunc("/login", s.handler.Users().Login()).Methods("POST")
	s.router.HandleFunc("/update-tokens", s.handler.Users().RefreshTokens()).Methods("POST")

	//news URLs
	s.router.HandleFunc("/news", s.handler.News().Create()).Methods("POST")
	s.router.HandleFunc("/news/{id}", s.handler.News().Get()).Methods("GET")
	s.router.HandleFunc("/news/{id}", s.handler.News().Change()).Methods("PUT")
	s.router.HandleFunc("/news/{id}", s.handler.News().Delete()).Methods("DELETE")

	//calendar URLs
	s.router.HandleFunc("/create-training", s.handler.Calendar().CreateWeek()).Methods("POST")
	s.router.HandleFunc("/get-calendar", s.handler.Calendar().GetWeek()).Methods("GET")
	s.router.HandleFunc("/training/{day}", s.handler.Calendar().ChangeDay()).Methods("PUT")

	//server middleware
	s.router.Use(s.LogRequest)
	s.router.Use(s.JwtAuthentication)
}
