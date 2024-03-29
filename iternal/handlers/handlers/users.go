package handlers

import (
	"CIS_Backend_Server/iternal/dto"
	"CIS_Backend_Server/iternal/handlers/response"
	"CIS_Backend_Server/iternal/model"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HandlerUser struct {
	handler *Handlers
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

		//checking that the structure and data type match in the request body
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

		err := h.handler.service.Users().CreateUser(a, u)

		//email and/or password not validated
		if errors.Is(err, model.ErrEmailPasswordNotValid) {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		//email already exists
		if errors.Is(err, model.ErrEmailIsBusy) {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		//server unexpected error
		if err != nil {
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

func (h *HandlerUser) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		//checking for a positive id value
		if id < 1 {
			response.Error(w, http.StatusBadRequest, model.ErrIDLessOne)
			return
		}

		//server unexpected route variable error
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		user, err := h.handler.service.Users().GetUser(id)

		//the user is not in the database
		if errors.Is(err, model.ErrUserNotFound) {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		//server unexpected error
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		response.Respond(w, http.StatusOK, user)
	}
}

func (h *HandlerUser) Login() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}

		//checking that the structure and data type match in the request body
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
			//email and/or password not validated
			if errors.Is(err, model.ErrEmailPasswordNotValid) {
				response.Error(w, http.StatusBadRequest, err)
				return
			}

			//login and/or password does not exist
			if errors.Is(err, model.ErrWrongEmailOrPassword) {
				response.Error(w, http.StatusBadRequest, err)
				return
			}

			//server unexpected error
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

		//checking that the structure and data type match in the request body
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		t := &model.Tokens{
			RefreshToken: req.RefreshToken,
		}

		if err := h.handler.service.Users().UpdateTokens(t); err != nil {
			//invalid token sent
			if errors.Is(err, model.ErrWrongToken) {
				response.Error(w, http.StatusBadRequest, err)
				return
			}

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
