package service

import (
	"CIS_Backend_Server/iternal/model"
	"errors"
	"fmt"
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
		if _, err := calendar[day[index]]; err == false {
			return &model.Err{Status: http.StatusBadRequest, Error: model.ErrIncorrectDaysOfWeek}
		}
	}

	//checking for an excess of the number of characters in the description of each day
	for _, key := range day {
		fmt.Println(len([]rune(calendar[key])))
		if len([]rune(calendar[key])) > 50 {
			return &model.Err{Status: http.StatusBadRequest, Error: model.ErrCharacterLimitExceeded}
		}
	}

	//send calendar to storage
	err := s.service.storage.Calendar().CreateWeek(calendar)

	//error check week already created
	if errors.Is(err, model.ErrWeekAlreadyCreated) {
		return &model.Err{Status: http.StatusBadRequest, Error: err}
	}

	return &model.Err{Status: http.StatusUnprocessableEntity, Error: err}
}

func (s *CalendarService) GetWeek() (trainings []model.Calendar, err error) {
	//query calendar from storage
	return s.service.storage.Calendar().GetWeek()
}

func (s *CalendarService) ChangeDay(calendar *model.Calendar) *model.Err {
	var err error

	//check for exceeding the number of characters in the description
	fmt.Println(len(calendar.Description), "|", len([]rune(calendar.Description)))
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
	if errors.Is(err, model.ErrWeekNotCreated) {
		return &model.Err{Status: http.StatusBadRequest, Error: err}
	}

	return &model.Err{Status: http.StatusUnprocessableEntity, Error: err}
}
