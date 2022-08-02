package handlers

import (
	"CIS_Backend_Server/iternal/dto"
	"CIS_Backend_Server/iternal/handlers/response"
	"CIS_Backend_Server/iternal/model"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HandlerNews struct {
	handler *Handlers
}

func (h *HandlerNews) Create() http.HandlerFunc {
	type request struct {
		Title       string
		Description string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		req.Title = r.FormValue("title")
		req.Description = r.FormValue("description")

		src, hdr, err := r.FormFile("photo")
		if err != nil {
			response.Error(w, http.StatusBadRequest, errors.New("wrong photo format"))
			return
		}
		object := &model.Photo{
			Payload: src,
			Size:    hdr.Size,
		}
		defer src.Close()

		n := &model.News{
			Title:       req.Title,
			Description: req.Description,
			Photo:       *object,
		}
		if err := h.handler.service.News().Create(r.Context(), n); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		data := &dto.News{
			Id:          n.Id,
			Title:       n.Title,
			Description: n.Description,
			Photo:       n.Name,
			TimeDate:    n.TimeDate,
		}
		response.Respond(w, http.StatusCreated, data)
	}
}

func (h *HandlerNews) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.handler.service.News().Get(r.Context())
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		response.Respond(w, http.StatusOK, data)
	}
}

func (h *HandlerNews) Change() http.HandlerFunc {
	type request struct {
		Id          int
		Title       string
		Description string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			response.Error(w, http.StatusNotFound, err)
			return
		}

		req.Title = r.FormValue("title")
		req.Description = r.FormValue("description")

		src, hdr, err := r.FormFile("photo")
		if err != nil {
			response.Error(w, http.StatusBadRequest, errors.New("wrong photo format"))
			return
		}
		object := &model.Photo{
			Payload: src,
			Size:    hdr.Size,
		}
		defer src.Close()

		n := &model.News{
			Id:          id,
			Title:       req.Title,
			Description: req.Description,
			Photo:       *object,
		}

		if err := h.handler.service.News().Change(r.Context(), n); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprintf("{News from id: %d has been successfully changed}", id)
		response.Respond(w, http.StatusOK, data)
	}
}

func (h *HandlerNews) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			response.Error(w, http.StatusNotFound, err)
			return
		}
		if err := h.handler.service.News().Delete(r.Context(), id); err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprintf("{News from id: %d was successfully deleted}", id)
		response.Respond(w, http.StatusOK, data)
	}
}
