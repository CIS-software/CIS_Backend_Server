package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"github.com/sirupsen/logrus"
)

type NewsRepository struct {
	storage *Storage
}

func (r *NewsRepository) CreateNews(e *model.News) error {
	return r.storage.db.QueryRow(
		"INSERT INTO news (title, description, photo) VALUES ($1, $2, $3) RETURNING id",
		e.Title,
		e.Description,
		e.Photo,
	).Scan(&e.Id)
}

func (r *NewsRepository) GetNews() (news []model.News, err error) {
	rows, err := r.storage.db.Query("SELECT * FROM news")
	if err != nil {
		logrus.Panic(err)
		//panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		e := model.News{}
		err := rows.Scan(&e.Id, &e.Title, &e.Description, &e.Photo)
		if err != nil {
			logrus.Info(err)
			continue
		}
		news = append(news, e)
	}
	return news, err
}
