package apiserver

import (
	"CIS_Backend_Server/iternal/app/apiserver/entities/handlers"
	"CIS_Backend_Server/iternal/app/apiserver/utils"
	"CIS_Backend_Server/iternal/app/model"
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type Server struct {
	router    *mux.Router
	logger    *logrus.Logger
	handler   *handlers.Handlers
	secretKey string
}

func newServer(logger *logrus.Logger, router *mux.Router, handler *handlers.Handlers, secretKey string) *Server {

	s := &Server{
		router:    router,
		logger:    logger,
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
	s.router.HandleFunc("/user/{id}", s.handler.Users().HandleGetUser()).Methods("GET")
	s.router.HandleFunc("/create-user", s.handler.Users().HandleCreateUser()).Methods("POST")
	s.router.HandleFunc("/login", s.handler.Users().HandleLogin()).Methods("POST")
	s.router.HandleFunc("/update-tokens", s.handler.Users().HandleUpdateTokens()).Methods("POST")
	s.router.HandleFunc("/news", s.handler.News().HandleCreateNews()).Methods("POST")
	s.router.HandleFunc("/news", s.handler.News().HandleGetNews()).Methods("GET")
	s.router.HandleFunc("/news/{id}", s.handler.News().HandleUpdateNews()).Methods("PUT")
	s.router.HandleFunc("/news/{id}", s.handler.News().HandleDeleteNews()).Methods("DELETE")
	s.router.HandleFunc("/create-training", s.handler.Calendar().HandleCreateTraining()).Methods("POST")
	s.router.HandleFunc("/get-calendar", s.handler.Calendar().HandleGetTrainingCalendar()).Methods("GET")
	s.router.HandleFunc("/training/{id}", s.handler.Calendar().HandleUpdateTraining()).Methods("PUT")
	s.router.HandleFunc("/training/{id}", s.handler.Calendar().HandleDeleteTraining()).Methods("DELETE")
	s.router.Use(s.JwtAuthentication)
}

func (s *Server) JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/login", "/create-user", "/update-tokens"}
		requestPath := r.URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("missing auth token"))
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("invalid auth token, does not match the format: Bearer {token-body}"))
			return
		}

		tokenPart := splitted[1]
		t := &model.Tokens{}
		token, err := jwt.ParseWithClaims(tokenPart, t, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secretKey), nil
		})
		if err != nil {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("wrong token"))
			return
		}

		if !token.Valid {
			utils.Error(w, r, http.StatusUnauthorized, errors.New("token is not valid"))
			return
		}

		ctx := context.WithValue(r.Context(), "user", t.TokenId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
