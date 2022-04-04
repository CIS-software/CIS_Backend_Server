package handlers

import (
	"CIS_Backend_Server/iternal/app/apiserver/utils"
	"CIS_Backend_Server/iternal/app/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HandlerNews struct {
	handler *Handlers
}

func (h *HandlerNews) HandleCreateNews() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Photo       string `json:"photo"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}

		n := &model.News{
			Title:       req.Title,
			Description: req.Description,
			Photo:       req.Photo,
		}
		if err := h.handler.service.News().CreateNews(n); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		utils.Respond(w, r, http.StatusCreated, n)
	}
}

func (h *HandlerNews) HandleGetNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.handler.service.News().GetNews()
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *HandlerNews) HandleUpdateNews() http.HandlerFunc {
	type request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Photo       string `json:"photo"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}

		e := &model.News{
			Id:          id,
			Title:       req.Title,
			Description: req.Description,
			Photo:       req.Photo,
		}
		if err := h.handler.service.News().UpdateNews(e); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprintf("{News from id: %d has been successfully changed}", id)
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *HandlerNews) HandleDeleteNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		if err := h.handler.service.News().DeleteNews(id); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprintf("{News from id: %d was successfully deleted}", id)
		utils.Respond(w, r, http.StatusOK, data)
	}
}
