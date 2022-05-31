package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"errors"
	"github.com/sirupsen/logrus"
)

type CalendarRepository struct {
	storage *Storage
}

func (r *CalendarRepository) CreateTrainingWeek(calendar map[string]string) error {
	days := [7]string{"пн", "вт", "ср", "чт", "пт", "сб", "вс"}
	for index := range days {
		if calendar[days[index]] == "" {
			return errors.New("wrong day")
		}
	}

	_, err := r.storage.db.Query("INSERT INTO training_calendar (day, description) "+
		"VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8), ($9, $10), ($11, $12), ($13, $14)",
		days[0], calendar[days[0]],
		days[1], calendar[days[1]],
		days[2], calendar[days[2]],
		days[3], calendar[days[3]],
		days[4], calendar[days[4]],
		days[5], calendar[days[5]],
		days[6], calendar[days[6]],
	)
	return err
}

func (r *CalendarRepository) GetTrainings() (calendar []model.Calendar, err error) {
	rows, err := r.storage.db.Query("SELECT day, description FROM training_calendar ORDER BY day")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := model.Calendar{}
		err := rows.Scan(&c.Day, &c.Description)
		if err != nil {
			logrus.Error(err)
			continue
		}
		calendar = append(calendar, c)
	}
	return calendar, err
}

func (r *CalendarRepository) UpdateTrainings(calendar *model.Calendar) error {
	result, err := r.storage.db.Exec("UPDATE training_calendar SET day = $1, description = $2 WHERE day = $3",
		calendar.Day,
		calendar.Description,
		calendar.Day,
	)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return errors.New("training not found")
	}
	return err
}
