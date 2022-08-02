package handlers

import (
	"CIS_Backend_Server/iternal/dto"
	"CIS_Backend_Server/iternal/handlers/response"
	"CIS_Backend_Server/iternal/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HandlerUser struct {
	handler *Handlers
}

func (h *HandlerUser) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if id < 1 {
			err = errors.New("id can't be negative")
			response.Error(w, http.StatusNotFound, err)
			return
		} else if err != nil {
			response.Error(w, http.StatusNotFound, err)
			return
		}
		user, err := h.handler.service.Users().GetUser(id)
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		response.Respond(w, http.StatusOK, user)
	}
}

func (h *HandlerUser) Create() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Name     string `json:"name"`
		Surname  string `json:"surname"`
		Town     string `json:"town"`
		Age      string `json:"age"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := new(request)
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		a := &model.UserAuth{
			Email:    req.Email,
			Password: req.Password,
		}
		u := &model.User{
			Name:    req.Name,
			Surname: req.Surname,
			Town:    req.Town,
			Age:     req.Age,
		}
		if err := h.handler.service.Users().CreateUser(a, u); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		data := &dto.User{
			Id:      a.Id,
			Email:   a.Email,
			Name:    u.Name,
			Surname: u.Surname,
			Town:    u.Town,
			Age:     u.Age,
		}

		response.Respond(w, http.StatusCreated, data)
	}
}

func (h *HandlerUser) Login() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		a := &model.UserAuth{
			Email:    req.Email,
			Password: req.Password,
		}
		t := new(model.Tokens)
		if err := h.handler.service.Users().Login(a, t); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		data := &model.Tokens{
			TokenId:      t.TokenId,
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
		}
		response.Respond(w, http.StatusCreated, data)
	}
}

func (h *HandlerUser) RefreshTokens() http.HandlerFunc {
	type request struct {
		RefreshToken string `json:"refresh-token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		t := &model.Tokens{
			RefreshToken: req.RefreshToken,
		}
		if err := h.handler.service.Users().UpdateTokens(t); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		data := &model.Tokens{
			TokenId:      t.TokenId,
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
		}
		response.Respond(w, http.StatusCreated, data)
	}
}
