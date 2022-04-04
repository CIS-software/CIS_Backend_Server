package service

import "CIS_Backend_Server/iternal/app/model"

type NewsService struct {
	service *Service
}

func (s *NewsService) CreateNews(n *model.News) error {
	return s.service.storage.News().CreateNews(n)
}

func (s *NewsService) GetNews() ([]model.News, error) {
	return s.service.storage.News().GetNews()
}

func (s *NewsService) UpdateNews(n *model.News) error {
	return s.service.storage.News().UpdateNews(n)
}

func (s *NewsService) DeleteNews(id int) error {
	return s.service.storage.News().DeleteNews(id)
}
