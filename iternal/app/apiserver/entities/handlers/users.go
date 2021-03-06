package handlers

import (
	"CIS_Backend_Server/iternal/app/apiserver/utils"
	"CIS_Backend_Server/iternal/app/dto"
	"CIS_Backend_Server/iternal/app/model"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HandlerUser struct {
	handler *Handlers
}

func (h *HandlerUser) HandleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if id < 1 {
			err = errors.New("id can't be negative")
			utils.Error(w, r, http.StatusNotFound, err)
			return
		} else if err != nil {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		user, err := h.handler.service.Users().GetUser(id)
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, user)
	}
}

func (h *HandlerUser) HandleCreateUser() http.HandlerFunc {
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
			utils.Error(w, r, http.StatusBadRequest, err)
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
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
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

		utils.Respond(w, r, http.StatusCreated, data)
	}
}

func (h *HandlerUser) HandleLogin() http.HandlerFunc {
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}

		a := &model.UserAuth{
			Email:    req.Email,
			Password: req.Password,
		}
		t := new(model.Tokens)
		if err := h.handler.service.Users().Login(a, t); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := &model.Tokens{
			TokenId:      t.TokenId,
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
		}
		utils.Respond(w, r, http.StatusCreated, data)
	}
}

func (h *HandlerUser) HandleUpdateTokens() http.HandlerFunc {
	type request struct {
		RefreshToken string `json:"refresh-token"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}

		t := &model.Tokens{
			RefreshToken: req.RefreshToken,
		}
		if err := h.handler.service.Users().UpdateTokens(t); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := &model.Tokens{
			TokenId:      t.TokenId,
			AccessToken:  t.AccessToken,
			RefreshToken: t.RefreshToken,
		}
		utils.Respond(w, r, http.StatusCreated, data)
	}
}
