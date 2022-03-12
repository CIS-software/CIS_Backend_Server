package storage

import "CIS_Backend_Server/iternal/app/model"

type NewsRepository interface {
	CreateNews(news *model.News) error
	GetNews() ([]model.News, error)
	UpdateNews(news *model.News) error
	DeleteNews(id int) error
}

type UsersRepository interface {
	CreateUser(user *model.User) error
	GetUser(id int) (users *model.User, err error)
	CreateUserAuth(user *model.UserAuth) error
	Login(user *model.UserAuth) error
	UpdateTokens(user *model.UserAuth) error
	GetUsers(id int) (users *model.User, err error)
}
