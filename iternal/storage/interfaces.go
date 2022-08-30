package storage

import (
	"CIS_Backend_Server/iternal/model"
	"context"
)

type Storage interface {
	News() NewsRepository
	Users() UsersRepository
	Calendar() CalendarRepository
}

type NewsRepository interface {
	Create(ctx context.Context, news *model.News) error
	Get(ctx context.Context, id int) ([]model.News, error)
	Change(ctx context.Context, news *model.News) error
	Delete(ctx context.Context, id int) error
}

type UsersRepository interface {
	CreateUser(userAuth *model.UserAuth, user *model.User) error
	GetUser(id int) (users *model.User, err error)
	Login(userAuth *model.UserAuth, tokens *model.Tokens) error
	UpdateTokens(tokens *model.Tokens) error
}

type CalendarRepository interface {
	CreateWeek(calendar map[string]string) error
	GetWeek() (trainings []model.Calendar, err error)
	ChangeDay(calendar *model.Calendar) error
}
