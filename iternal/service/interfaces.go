package service

import (
	"CIS_Backend_Server/iternal/model"
	"context"
)

type Service interface {
	News() NewsService
	Users() UsersService
	Calendar() CalendarService
}

type NewsService interface {
	Create(ctx context.Context, news *model.News) error
	Get(ctx context.Context) ([]model.News, error)
	Change(ctx context.Context, news *model.News) error
	Delete(ctx context.Context, id int) error
}

type UsersService interface {
	CreateUser(userAuth *model.UserAuth, user *model.User) error
	GetUser(id int) (users *model.User, err error)
	Login(userAuth *model.UserAuth, tokens *model.Tokens) error
	UpdateTokens(tokens *model.Tokens) error
}

type CalendarService interface {
	CreateWeek(calendar map[string]string) *model.Err
	GetWeek() (trainings []model.Calendar, err error)
	ChangeDay(calendar *model.Calendar) error
}
