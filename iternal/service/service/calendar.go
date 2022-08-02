package service

import (
	"CIS_Backend_Server/iternal/model"
	"errors"
	"net/http"
)

var (
	ErrNoEqualSeven           = errors.New("wrong number of days passed, expected 7")
	ErrCharacterLimitExceeded = errors.New("description cannot be more than 50 characters")
	ErrNotEnoughCharacters    = errors.New("description must be at least 2 characters long")
	ErrIncorrectDaysOfWeek    = errors.New("incorrect names of the days of the week")
)

type CalendarService struct {
	service *Service
}

func (s *CalendarService) CreateWeek(calendar map[string]string) *model.Err {
	var day = [7]string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}

	//Checking the size of the map, the size of the map is the number of days in a week, i.e. 7
	if len(calendar) != 7 {
		return &model.Err{Status: http.StatusBadRequest, Error: ErrNoEqualSeven}
	}

	//Checking the correctness of the names of the days of the week and their uniqueness
	for index := range day {
		var dayOfWeek = 0
		for key := range calendar {
			dayOfWeek++
			if key == day[index] {
				break
			}
			if dayOfWeek == 7 {
				return &model.Err{Status: http.StatusBadRequest, Error: ErrIncorrectDaysOfWeek}
			}
		}
	}

	//Checking the description of each day for the min and max number of characters
	for _, description := range calendar {
		if len(description) < 2 {
			return &model.Err{Status: http.StatusBadRequest, Error: ErrNotEnoughCharacters}
		}
		if len(description) > 50 {
			return &model.Err{Status: http.StatusBadRequest, Error: ErrCharacterLimitExceeded}
		}
	}

	err := s.service.storage.Calendar().CreateWeek(calendar)
	return &model.Err{Status: http.StatusUnprocessableEntity, Error: err}
}

func (s *CalendarService) GetWeek() (trainings []model.Calendar, err error) {
	return s.service.storage.Calendar().GetWeek()
}

func (s *CalendarService) ChangeDay(calendar *model.Calendar) error {
	return s.service.storage.Calendar().ChangeDay(calendar)
}
