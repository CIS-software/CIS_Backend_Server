package apiserver

import (
	"CIS_Backend_Server/iternal/app/model"
	"CIS_Backend_Server/iternal/app/storage"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

type server struct {
	router    *mux.Router
	logger    *logrus.Logger
	storage   storage.Storage
	secretKey string
}

func newServer(storage storage.Storage, logger *logrus.Logger, secretKey string) *server {

	s := &server{
		router:    mux.NewRouter(),
		logger:    logger,
		storage:   storage,
		secretKey: secretKey,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/user/{id}", s.handleGetUser()).Methods("POST")
	s.router.HandleFunc("/user", s.handleCreateUser()).Methods("POST")
	s.router.HandleFunc("/create-user-auth", s.handleCreateUserAuth()).Methods("POST")
	s.router.HandleFunc("/login", s.handleLogin()).Methods("POST")
	s.router.HandleFunc("/update-tokens", s.handleUpdateTokens()).Methods("POST")
	s.router.HandleFunc("/news", s.handleCreateNews()).Methods("POST")
	s.router.HandleFunc("/news", s.handleGetNews()).Methods("GET")
	s.router.HandleFunc("/news/{id}", s.handleUpdateNews()).Methods("PUT")
	s.router.HandleFunc("/news/{id}", s.handleDeleteNews()).Methods("DELETE")
	s.router.Use(s.JwtAuthentication)
}

func (s *server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if id < 1 {
			err = errors.New("id can't be negative")
			s.error(w, r, http.StatusNotFound, err)
			return
		} else if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		user, err := s.storage.Users().GetUser(id)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, user)
	}
}

func (s *server) handleCreateUser() http.HandlerFunc {
	type request struct {
		Name       string `json:"name"`
		Surname    string `json:"surname"`
		Patronymic string `json:"patronymic"`
		Town       string `json:"town"`
		Age        int    `json:"age"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		userId, err := strconv.Atoi(fmt.Sprintf("%v", r.Context().Value("user")))
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
		e := &model.User{
			UserID:     userId,
			Name:       req.Name,
			Surname:    req.Surname,
			Patronymic: req.Patronymic,
			Town:       req.Town,
			Age:        req.Age,
		}
		if err := s.storage.Users().CreateUser(e); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, e)
	}
}

func (s *server) handleCreateUserAuth() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.UserAuth{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.storage.Users().CreateUserAuth(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) handleLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.UserAuth{
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.storage.Users().Login(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		data := &model.Tokens{
			AccessToken:  u.AccessToken,
			RefreshToken: u.RefreshToken,
		}
		s.respond(w, r, http.StatusCreated, data)
	}
}

func (s *server) handleUpdateTokens() http.HandlerFunc {
	type request struct {
		RefreshToken string `json:"refresh-token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.UserAuth{
			RefreshToken: req.RefreshToken,
		}
		if err := s.storage.Users().UpdateTokens(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		data := &model.Tokens{
			AccessToken:  u.AccessToken,
			RefreshToken: u.RefreshToken,
		}
		s.respond(w, r, http.StatusCreated, data)
	}
}

func (s *server) handleCreateNews() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Photo       string `json:"photo"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.News{
			Title:       req.Title,
			Description: req.Description,
			Photo:       req.Photo,
		}
		if err := s.storage.News().CreateNews(e); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, e)
	}
}

func (s *server) handleGetNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.storage.News().GetNews()
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, data)
	}
}

func (s *server) handleUpdateNews() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Photo       string `json:"photo"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.News{
			Id:          id,
			Title:       req.Title,
			Description: req.Description,
			Photo:       req.Photo,
		}
		if err := s.storage.News().UpdateNews(e); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		result := fmt.Sprintf("{News from id: %d has been successfully changed}", id)
		s.respond(w, r, http.StatusOK, result)
	}
}

func (s *server) handleDeleteNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		if err := s.storage.News().DeleteNews(id); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		result := fmt.Sprintf("{News from id: %d was successfully deleted}", id)
		s.respond(w, r, http.StatusOK, result)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	s.logger.Info("Status: ", code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func (s *server) JwtAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		notAuth := []string{"/login", "/create-user-auth", "/update-tokens"}
		requestPath := r.URL.Path
		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			s.error(w, r, http.StatusUnauthorized, errors.New("missing auth token"))
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			s.error(w, r, http.StatusUnauthorized, errors.New("invalid auth token, does not match the format: Bearer {token-body}"))
			return
		}

		tokenPart := splitted[1]
		u := &model.UserAuth{}
		token, err := jwt.ParseWithClaims(tokenPart, u, func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secretKey), nil
		})
		if err != nil {
			s.error(w, r, http.StatusUnauthorized, errors.New("wrong token"))
			return
		}

		if !token.Valid {
			s.error(w, r, http.StatusUnauthorized, errors.New("token is not valid"))
			return
		}

		ctx := context.WithValue(r.Context(), "user", u.Id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
