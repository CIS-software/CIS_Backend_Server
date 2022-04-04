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

type HandlerCalendar struct {
	handler *Handlers
}

func (h *HandlerCalendar) HandleCreateTraining() http.HandlerFunc {
	type request struct {
		Date        string `json:"date"`
		Description string `json:"description"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			utils.Error(w, r, http.StatusBadRequest, err)
			return
		}

		c := &model.Calendar{
			Date:        req.Date,
			Description: req.Description,
		}
		if err := h.handler.service.Calendar().CreateTraining(c); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		utils.Respond(w, r, http.StatusCreated, c)
	}
}

func (h *HandlerCalendar) HandleGetTrainingCalendar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := h.handler.service.Calendar().GetTrainings()
		if err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *HandlerCalendar) HandleUpdateTraining() http.HandlerFunc {
	type request struct {
		Date        string `json:"date"`
		Description string `json:"description"`
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

		c := &model.Calendar{
			Id:          id,
			Date:        req.Date,
			Description: req.Description,
		}
		if err := h.handler.service.Calendar().UpdateTrainings(c); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprintf("{Training from id: %d has been successfully changed}", id)
		utils.Respond(w, r, http.StatusOK, data)
	}
}

func (h *HandlerCalendar) HandleDeleteTraining() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil || id < 1 {
			utils.Error(w, r, http.StatusNotFound, err)
			return
		}
		if err := h.handler.service.Calendar().DeleteTrainings(id); err != nil {
			utils.Error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		data := fmt.Sprintf("{Training from id: %d was successfully deleted}", id)
		utils.Respond(w, r, http.StatusOK, data)
	}
}
