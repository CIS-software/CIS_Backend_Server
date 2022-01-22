package storage

import "CIS_Backend_Server/iternal/app/model"

type NewsRepository interface {
	CreateNews(news *model.News) error
	GetNews() ([]model.News, error)
}
