package service

import (
	"CIS_Backend_Server/iternal/app/model"
	"context"
)

type NewsService struct {
	service *Service
}

func (s *NewsService) CreateNews(ctx context.Context, n *model.News) error {
	return s.service.storage.News().CreateNews(ctx, n)
}

func (s *NewsService) GetNews(ctx context.Context) ([]model.News, error) {
	return s.service.storage.News().GetNews(ctx)
}

func (s *NewsService) UpdateNews(ctx context.Context, n *model.News) error {
	return s.service.storage.News().UpdateNews(ctx, n)
}

func (s *NewsService) DeleteNews(ctx context.Context, id int) error {
	return s.service.storage.News().DeleteNews(ctx, id)
}
