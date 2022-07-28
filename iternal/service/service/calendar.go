package service

import (
	"CIS_Backend_Server/iternal/model"
)

type CalendarService struct {
	service *Service
}

func (s *CalendarService) CreateWeek(calendar map[string]string) error {
	return s.service.storage.Calendar().CreateWeek(calendar)
}

func (s *CalendarService) GetWeek() (trainings []model.Calendar, err error) {
	return s.service.storage.Calendar().GetWeek()
}

func (s *CalendarService) ChangeDay(calendar *model.Calendar) error {
	return s.service.storage.Calendar().ChangeDay(calendar)
}
