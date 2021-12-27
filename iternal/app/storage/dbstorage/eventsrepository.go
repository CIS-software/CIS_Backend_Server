package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"github.com/sirupsen/logrus"
)

type EventsRepository struct {
	storage *Storage
}

func (r *EventsRepository) CreateEvents(e *model.Events) error {
	return r.storage.db.QueryRow(
		"INSERT INTO events (title, description, photo) VALUES ($1, $2, $3) RETURNING id",
		e.Title,
		e.Description,
		e.Photo,
	).Scan(&e.Id)
}

func (r *EventsRepository) GetEvents() (events []model.Events, err error){
	rows, err := r.storage.db.Query("SELECT * FROM events")
	if err != nil {
		logrus.Panic(err)
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		e := model.Events{}
		err := rows.Scan(&e.Id, &e.Title, &e.Description, &e.Photo)
		if err != nil{
			logrus.Info(err)
			continue
		}
		events = append(events, e)
	}
	return events, err
}