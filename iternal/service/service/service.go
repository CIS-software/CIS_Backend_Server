package service

import (
	"CIS_Backend_Server/iternal/service"
	"CIS_Backend_Server/iternal/storage"
)

type Service struct {
	storage         storage.Storage
	newsService     *NewsService
	userService     *UsersService
	calendarService *CalendarService
}

func New(storage storage.Storage) *Service {
	return &Service{storage: storage}
}

func (s *Service) News() service.NewsService {
	if s.newsService != nil {
		return s.newsService
	}

	s.newsService = &NewsService{
		service: s,
	}
	return s.newsService
}

func (s *Service) Users() service.UsersService {
	if s.userService != nil {
		return s.userService
	}

	s.userService = &UsersService{
		service: s,
	}
	return s.userService
}

func (s *Service) Calendar() service.CalendarService {
	if s.calendarService != nil {
		return s.calendarService
	}

	s.calendarService = &CalendarService{
		service: s,
	}
	return s.calendarService
}
