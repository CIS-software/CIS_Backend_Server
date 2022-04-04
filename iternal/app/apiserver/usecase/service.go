package usecase

import (
	"CIS_Backend_Server/iternal/app/model"
)

type Service interface {
	News() NewsService
	Users() UsersService
	Calendar() CalendarService
}

type NewsService interface {
	CreateNews(news *model.News) error
	GetNews() ([]model.News, error)
	UpdateNews(news *model.News) error
	DeleteNews(id int) error
}

type UsersService interface {
	CreateUser(userAuth *model.UserAuth, user *model.User) error
	GetUser(id int) (users *model.User, err error)
	Login(userAuth *model.UserAuth, tokens *model.Tokens) error
	UpdateTokens(tokens *model.Tokens) error
}

type CalendarService interface {
	CreateTraining(calendar *model.Calendar) error
	GetTrainings() (trainings []model.Calendar, err error)
	UpdateTrainings(calendar *model.Calendar) error
	DeleteTrainings(id int) error
}
