package service

import (
	"CIS_Backend_Server/iternal/model"
	"context"
)

type NewsService struct {
	service *Service
}

func (s *NewsService) Create(ctx context.Context, n *model.News) error {
	return s.service.storage.News().Create(ctx, n)
}

func (s *NewsService) Get(ctx context.Context) ([]model.News, error) {
	return s.service.storage.News().Get(ctx)
}

func (s *NewsService) Change(ctx context.Context, n *model.News) error {
	return s.service.storage.News().Change(ctx, n)
}

func (s *NewsService) Delete(ctx context.Context, id int) error {
	return s.service.storage.News().Delete(ctx, id)
}
