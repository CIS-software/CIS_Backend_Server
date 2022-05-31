package usecase

import (
	"CIS_Backend_Server/iternal/app/model"
	"context"
)

type Service interface {
	News() NewsService
	Users() UsersService
	Calendar() CalendarService
}

type NewsService interface {
	CreateNews(ctx context.Context, news *model.News) error
	GetNews(ctx context.Context) ([]model.News, error)
	UpdateNews(ctx context.Context, news *model.News) error
	DeleteNews(ctx context.Context, id int) error
}

type UsersService interface {
	CreateUser(userAuth *model.UserAuth, user *model.User) error
	GetUser(id int) (users *model.User, err error)
	Login(userAuth *model.UserAuth, tokens *model.Tokens) error
	UpdateTokens(tokens *model.Tokens) error
}

type CalendarService interface {
	CreateTrainingWeek(calendar map[string]string) error
	GetTrainings() (trainings []model.Calendar, err error)
	UpdateTrainings(calendar *model.Calendar) error
}
