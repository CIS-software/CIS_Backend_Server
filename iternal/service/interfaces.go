package service

import (
	model2 "CIS_Backend_Server/iternal/model"
	"context"
)

type Service interface {
	News() NewsService
	Users() UsersService
	Calendar() CalendarService
}

type NewsService interface {
	Create(ctx context.Context, news *model2.News) error
	Get(ctx context.Context) ([]model2.News, error)
	Change(ctx context.Context, news *model2.News) error
	Delete(ctx context.Context, id int) error
}

type UsersService interface {
	CreateUser(userAuth *model2.UserAuth, user *model2.User) error
	GetUser(id int) (users *model2.User, err error)
	Login(userAuth *model2.UserAuth, tokens *model2.Tokens) error
	UpdateTokens(tokens *model2.Tokens) error
}

type CalendarService interface {
	CreateWeek(calendar map[string]string) error
	GetWeek() (trainings []model2.Calendar, err error)
	ChangeDay(calendar *model2.Calendar) error
}
