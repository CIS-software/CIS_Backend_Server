package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type NewsRepository struct {
	storage *Storage
}

func (r *NewsRepository) CreateNews(e *model.News) error {
	return r.storage.db.QueryRow(
		"INSERT INTO news (title, description, photo) VALUES ($1, $2, $3) RETURNING id, time_date",
		e.Title,
		e.Description,
		e.Photo,
	).Scan(&e.Id, &e.TimeDate)
}

func (r *NewsRepository) GetNews() (news []model.News, err error) {
	rows, err := r.storage.db.Query("SELECT * FROM news ORDER BY time_date DESC")
	if err != nil {
		logrus.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		e := model.News{}
		err := rows.Scan(&e.Id, &e.Title, &e.Description, &e.Photo, &e.TimeDate)
		if err != nil {
			logrus.Info(err)
			continue
		}
		news = append(news, e)
	}
	return news, err
}

func (r *NewsRepository) UpdateNews(e *model.News) error {
	result, err := r.storage.db.Exec("UPDATE news SET title = $1, description = $2, photo = $3 WHERE id = $4",
		e.Title,
		e.Description,
		e.Photo,
		e.Id,
	)
	if err != nil {
		logrus.Panic(err)
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("news not found")
	}
	return err
}

func (r *NewsRepository) DeleteNews(id int) error {
	result, err := r.storage.db.Exec("DELETE FROM news WHERE id = $1", id)
	if err != nil {
		logrus.Panic(err)
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("news not found")
	}
	return err
}
