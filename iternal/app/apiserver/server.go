package apiserver

import (
	"CIS_Backend_Server/iternal/app/model"
	"CIS_Backend_Server/iternal/app/storage"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type server struct {
	router  *mux.Router
	logger  *logrus.Logger
	storage storage.Storage
}

func newServer(storage storage.Storage, logger *logrus.Logger) *server {
	s := &server{
		router:  mux.NewRouter(),
		logger:  logger,
		storage: storage,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users/{id}", s.handleGetUser()).Methods("GET")
	s.router.HandleFunc("/users", s.handleCreateUser()).Methods("POST")
	s.router.HandleFunc("/news", s.handleCreateNews()).Methods("POST")
	s.router.HandleFunc("/news", s.handleGetNews()).Methods("GET")
	s.router.HandleFunc("/news/{id}", s.handleUpdateNews()).Methods("PUT")
	s.router.HandleFunc("/news/{id}", s.handleDeleteNews()).Methods("DELETE")
}

func (s *server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if id < 1 {
			err = errors.New("Id can't be negative")
			s.error(w, r, http.StatusNotFound, err)
			return
		} else if err != nil {
			s.error(w, r, http.StatusNotFound, err)
			return
		}
		user, err := s.storage.Users().GetUsers(id)
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

		e := &model.User{
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
