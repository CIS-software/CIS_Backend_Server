package service

import (
	"CIS_Backend_Server/iternal/app/apiserver/usecase"
	"CIS_Backend_Server/iternal/app/storage"
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

func (s *Service) News() usecase.NewsService {
	if s.newsService != nil {
		return s.newsService
	}

	s.newsService = &NewsService{
		service: s,
	}
	return s.newsService
}

func (s *Service) Users() usecase.UsersService {
	if s.userService != nil {
		return s.userService
	}

	s.userService = &UsersService{
		service: s,
	}
	return s.userService
}

func (s *Service) Calendar() usecase.CalendarService {
	if s.calendarService != nil {
		return s.calendarService
	}

	s.calendarService = &CalendarService{
		service: s,
	}
	return s.calendarService
}
