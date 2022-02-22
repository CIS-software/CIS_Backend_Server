package storage

import "CIS_Backend_Server/iternal/app/model"

type NewsRepository interface {
	CreateNews(news *model.News) error
	GetNews() ([]model.News, error)
	UpdateNews(news *model.News) error
	DeleteNews(news *model.News) error
}

type UsersRepository interface {
	CreateUser(user *model.User) error
	GetUsers() (users []model.User, err error)
	Login(user *model.User) (uint64, error)
}
