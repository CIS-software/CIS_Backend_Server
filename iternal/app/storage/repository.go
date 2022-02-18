package storage

import "CIS_Backend_Server/iternal/app/model"

type NewsRepository interface {
	CreateNews(news *model.News) error
	GetNews() ([]model.News, error)
	UpdateNews(news *model.News) error
	DeleteNews(news *model.News) error
}
