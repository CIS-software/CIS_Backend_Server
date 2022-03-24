package storage

import "CIS_Backend_Server/iternal/app/model"

type NewsRepository interface {
	CreateNews(news *model.News) error
	GetNews() ([]model.News, error)
	UpdateNews(news *model.News) error
	DeleteNews(id int) error
}

type UsersRepository interface {
	CreateUser(userAuth *model.UserAuth, user *model.User) error
	GetUser(id int) (users *model.User, err error)
	Login(userAuth *model.UserAuth, tokens *model.Tokens) error
	UpdateTokens(tokens *model.Tokens) error
}
