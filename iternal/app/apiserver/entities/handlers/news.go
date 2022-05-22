package handlers

import (
	"CIS_Backend_Server/iternal/app/apiserver/utils"
	"CIS_Backend_Server/iternal/app/dto"
	"CIS_Backend_Server/iternal/app/model"
	"errors"
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
		Title       string
		Description string
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		req.Title = r.FormValue("title")
		req.Description = r.FormValue("description")

		src, hdr, err := r.FormFile("photo")
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, errors.New("wrong photo format"))
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
		if err := h.handler.service.News().CreateNews(r.Context(), n); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := &dto.News{
			Id:          n.Id,
			Title:       n.Title,
			Description: n.Description,
			Photo:       n.Name,
			TimeDate:    n.TimeDate,
		}
		utils.Respond(w, r, http.StatusCreated, data)
	}
}

func (h *HandlerNews) HandleGetNews() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.handler.service.News().GetNews(r.Context())
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *HandlerNews) HandleUpdateNews() http.HandlerFunc {
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
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}

		req.Title = r.FormValue("title")
		req.Description = r.FormValue("description")

		src, hdr, err := r.FormFile("photo")
		if err != nil {
			utils.Error(w, r, http.StatusBadRequest, errors.New("wrong photo format"))
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

		if err := h.handler.service.News().UpdateNews(r.Context(), n); err != nil {
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
		if err := h.handler.service.News().DeleteNews(r.Context(), id); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprintf("{News from id: %d was successfully deleted}", id)
		utils.Respond(w, r, http.StatusOK, data)
	}
}
