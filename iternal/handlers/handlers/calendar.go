package handlers

import (
	"CIS_Backend_Server/iternal/model"
	"CIS_Backend_Server/iternal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

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
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}
		logrus.Info(req.TrainingCalendar, len(req.TrainingCalendar))
		if len(req.TrainingCalendar) != 7 {
			utils.Error(w, r, http.StatusBadRequest, errors.New("map size is not 7, expected 7 key-value pairs"))
			return
		}

		if err := h.handler.service.Calendar().CreateWeek(req.TrainingCalendar); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		data := fmt.Sprint("Training week successfully created")
		utils.Respond(w, r, http.StatusCreated, data)
	}
}

func (h *HandlerCalendar) GetWeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.handler.service.Calendar().GetWeek()
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
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
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}

		c := &model.Calendar{
			Day:         day,
			Description: req.Description,
		}
		if err := h.handler.service.Calendar().ChangeDay(c); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprint("{Training day successfully changed}")
		utils.Respond(w, r, http.StatusOK, data)
	}
}
