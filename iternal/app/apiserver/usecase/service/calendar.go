package service

import "CIS_Backend_Server/iternal/app/model"

type CalendarService struct {
	service *Service
}

func (s *CalendarService) CreateTrainingWeek(calendar map[string]string) error {
	return s.service.storage.Calendar().CreateTrainingWeek(calendar)
}

func (s *CalendarService) GetTrainings() (trainings []model.Calendar, err error) {
	return s.service.storage.Calendar().GetTrainings()
}

func (s *CalendarService) UpdateTrainings(calendar *model.Calendar) error {
	return s.service.storage.Calendar().UpdateTrainings(calendar)
}
