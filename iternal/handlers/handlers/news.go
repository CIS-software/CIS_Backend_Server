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

		//file existence check
		if errors.Is(err, http.ErrMissingFile) {
			response.Error(w, http.StatusBadRequest, http.ErrMissingFile)
			return
		}

		//server unexpected error
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		object := &model.Photo{
			Payload: src,
			Name:    hdr.Filename,
			Size:    hdr.Size,
		}
		defer src.Close()

		n := &model.News{
			Title:       req.Title,
			Description: req.Description,
			Photo:       *object,
		}

		if err := h.handler.service.News().Create(r.Context(), n); err != nil {
			//with wrong file extension
			if errors.Is(err, model.ErrWrongContentType) {
				response.Error(w, http.StatusBadRequest, model.ErrWrongContentType)
				return
			}

			//with a long file name
			if errors.Is(err, model.ErrLongFileName) {
				response.Error(w, http.StatusBadRequest, model.ErrLongFileName)
				return
			}

			//with not validated title and/or description
			if errors.Is(err, model.ErrNewsNotValid) {
				response.Error(w, http.StatusBadRequest, model.ErrNewsNotValid)
				return
			}

			//server unexpected error
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		//response body structure
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
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])

		//checking for a non-negative id value
		if id < 0 {
			response.Error(w, http.StatusBadRequest, model.ErrNegativeID)
			return
		}

		//server unexpected route variable error
		if err != nil {
			response.Error(w, http.StatusNotFound, err)
			return
		}

		data, err := h.handler.service.News().Get(r.Context(), id)
		//when there is no news in the database
		if errors.Is(err, model.ErrNewsNotFound) {
			response.Error(w, http.StatusBadRequest, model.ErrNewsNotFound)
			return
		}

		//server unexpected error
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

		//checking for a positive id value
		if id < 1 {
			response.Error(w, http.StatusBadRequest, model.ErrIDLessOne)
			return
		}

		//server unexpected route variable error
		if err != nil {
			response.Error(w, http.StatusNotFound, err)
			return
		}

		req.Title = r.FormValue("title")
		req.Description = r.FormValue("description")

		src, hdr, err := r.FormFile("photo")

		//file existence check
		if errors.Is(err, http.ErrMissingFile) {
			response.Error(w, http.StatusBadRequest, http.ErrMissingFile)
			return
		}

		//server unexpected error
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		object := &model.Photo{
			Payload: src,
			Name:    hdr.Filename,
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
			//with wrong file extension
			if errors.Is(err, model.ErrWrongContentType) {
				response.Error(w, http.StatusBadRequest, model.ErrWrongContentType)
				return
			}

			//with a long file name
			if errors.Is(err, model.ErrLongFileName) {
				response.Error(w, http.StatusBadRequest, model.ErrLongFileName)
				return
			}

			//with not validated title and/or description
			if errors.Is(err, model.ErrNewsNotValid) {
				response.Error(w, http.StatusBadRequest, model.ErrNewsNotValid)
				return
			}

			//when there is no news in the database
			if errors.Is(err, model.ErrNewsNotFound) {
				response.Error(w, http.StatusBadRequest, model.ErrNewsNotFound)
				return
			}

			//server unexpected error
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		data := fmt.Sprint("{News has been successfully changed}")

		response.Respond(w, http.StatusOK, data)
	}
}

func (h *HandlerNews) Delete() http.HandlerFunc {
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
			response.Error(w, http.StatusNotFound, err)
			return
		}

		if err := h.handler.service.News().Delete(r.Context(), id); err != nil {
			//when there is no news in the database
			if errors.Is(err, model.ErrNewsNotFound) {
				response.Error(w, http.StatusBadRequest, model.ErrNewsNotFound)
				return
			}

			//server unexpected error
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}

		data := fmt.Sprint("{News was successfully deleted}")

		response.Respond(w, http.StatusOK, data)
	}
}
