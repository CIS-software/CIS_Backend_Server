package handlers

import (
	"CIS_Backend_Server/iternal/handlers/response"
	"CIS_Backend_Server/iternal/model"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

//var (
//	ErrWeekNotCreated = errors.New("training week is not in the database")
//)

type HandlerCalendar struct {
	handler *Handlers
}

func (h *HandlerCalendar) CreateWeek() http.HandlerFunc {
	type request struct {
		TrainingCalendar map[string]string `json:"training-calendar"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		if err := h.handler.service.Calendar().CreateWeek(req.TrainingCalendar); err.Error != nil {
			if err.Error == model.ErrWeekAlreadyCreated {
				response.Error(w, http.StatusBadRequest, model.ErrWeekAlreadyCreated)
				return
			}
			response.Error(w, err.Status, err.Error)
			return
		}

		success := fmt.Sprint("Training week successfully created")
		response.Respond(w, http.StatusCreated, success)
	}
}

func (h *HandlerCalendar) GetWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		calendar, err := h.handler.service.Calendar().GetWeek()
		if err != nil {
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		if calendar == nil {
			response.Error(w, http.StatusBadRequest, model.ErrWeekNotCreated)
			return
		}
		response.Respond(w, http.StatusOK, calendar)
	}
}

func (h *HandlerCalendar) ChangeDay() http.HandlerFunc {
	type request struct {
		Description string `json:"description"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		day := vars["day"]
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			response.Error(w, http.StatusBadRequest, err)
			return
		}

		c := &model.Calendar{
			Day:         day,
			Description: req.Description,
		}
		if err := h.handler.service.Calendar().ChangeDay(c); err != nil {
			if errors.Is(err, model.ErrWeekNotCreated) {
				response.Error(w, http.StatusBadRequest, model.ErrWeekNotCreated)
				return
			}
			response.Error(w, http.StatusUnprocessableEntity, err)
			return
		}
		success := fmt.Sprint("Training day successfully changed")
		response.Respond(w, http.StatusOK, success)
	}
}
