package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type CalendarRepository struct {
	storage *Storage
}

func (r *CalendarRepository) CreateTraining(c *model.Calendar) error {
	return r.storage.db.QueryRow("INSERT INTO training_calendar (date, description) VALUES ($1, $2) RETURNING id",
		c.Date,
		c.Description,
	).Scan(&c.Id)
}

func (r *CalendarRepository) GetTrainings() (trainings []model.Calendar, err error) {
	rows, err := r.storage.db.Query("SELECT id, date, description FROM training_calendar")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := model.Calendar{}
		err := rows.Scan(&c.Id, &c.Date, &c.Description)
		if err != nil {
			logrus.Error(err)
			continue
		}
		trainings = append(trainings, c)
	}
	return trainings, err
}

func (r *CalendarRepository) UpdateTrainings(c *model.Calendar) error {
	result, err := r.storage.db.Exec("UPDATE training_calendar SET date = $1, description = $2 WHERE id = $3",
		c.Date,
		c.Description,
		c.Id,
	)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("training not found")
	}
	return err
}

func (r *CalendarRepository) DeleteTrainings(id int) error {
	result, err := r.storage.db.Exec("DELETE FROM training_calendar WHERE id = $1", id)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("training not found")
	}
	return err
}
