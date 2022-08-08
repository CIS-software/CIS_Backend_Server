package service

import (
	"CIS_Backend_Server/iternal/model"
	"errors"
	"net/http"
)

var (
	day = [7]string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}
)

type CalendarService struct {
	service *Service
}

func (s *CalendarService) CreateWeek(calendar map[string]string) *model.Err {
	//checking the size of the map, the size of the map is the number of days in a week, i.e. 7
	if len(calendar) != 7 {
		return &model.Err{Status: http.StatusBadRequest, Error: model.ErrNoEqualSeven}
	}

	//checking the correctness of the names of the days of the week and their uniqueness
	for index := range day {
		var dayOfWeek = 0
		for key := range calendar {
			dayOfWeek++
			if key == day[index] {
				break
			}
			if dayOfWeek == 7 {
				return &model.Err{Status: http.StatusBadRequest, Error: model.ErrIncorrectDaysOfWeek}
			}
		}
	}

	//checking the description of each day for the min and max number of characters
	for _, description := range calendar {
		if len(description) < 2 {
			return &model.Err{Status: http.StatusBadRequest, Error: model.ErrNotEnoughCharacters}
		}
		if len(description) > 50 {
			return &model.Err{Status: http.StatusBadRequest, Error: model.ErrCharacterLimitExceeded}
		}
	}

	err := s.service.storage.Calendar().CreateWeek(calendar)

	//error check week already created
	if errors.Is(err, model.ErrWeekAlreadyCreated) {
		return &model.Err{Status: http.StatusBadRequest, Error: err}
	}

	return &model.Err{Status: http.StatusUnprocessableEntity, Error: err}
}

func (s *CalendarService) GetWeek() (trainings []model.Calendar, err error) {
	return s.service.storage.Calendar().GetWeek()
}

func (s *CalendarService) ChangeDay(calendar *model.Calendar) *model.Err {
	var err error

	//checking the description of the day for the minimum and maximum number of characters
	if len(calendar.Description) < 2 {
		return &model.Err{Status: http.StatusBadRequest, Error: model.ErrNotEnoughCharacters}
	}
	if len(calendar.Description) > 50 {
		return &model.Err{Status: http.StatusBadRequest, Error: model.ErrCharacterLimitExceeded}
	}

	//day of the week check
	for index := range day {
		if calendar.Day == day[index] {
			err = s.service.storage.Calendar().ChangeDay(calendar)
			break
		}
		if index == 6 {
			return &model.Err{Status: http.StatusBadRequest, Error: model.ErrIncorrectDaysOfWeek}
		}
	}

	//checking for a table not created error
	if errors.Is(err, model.ErrWeekAlreadyCreated) {
		return &model.Err{Status: http.StatusBadRequest, Error: err}
	}

	return &model.Err{Status: http.StatusUnprocessableEntity, Error: err}
}
