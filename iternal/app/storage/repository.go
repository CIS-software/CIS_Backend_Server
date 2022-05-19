package storage

import (
	"CIS_Backend_Server/iternal/app/model"
	"context"
)

type NewsRepository interface {
	CreateNews(ctx context.Context, news *model.News) error
	GetNews(ctx context.Context) ([]model.News, error)
	UpdateNews(ctx context.Context, news *model.News) error
	DeleteNews(ctx context.Context, id int) error
}

type UsersRepository interface {
	CreateUser(userAuth *model.UserAuth, user *model.User) error
	GetUser(id int) (users *model.User, err error)
	Login(userAuth *model.UserAuth, tokens *model.Tokens) error
	UpdateTokens(tokens *model.Tokens) error
}

type CalendarRepository interface {
	CreateTraining(calendar *model.Calendar) error
	GetTrainings() (trainings []model.Calendar, err error)
	UpdateTrainings(calendar *model.Calendar) error
	DeleteTrainings(id int) error
}
