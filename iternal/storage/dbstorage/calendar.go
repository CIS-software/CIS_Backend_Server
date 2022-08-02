package dbstorage

import (
	"CIS_Backend_Server/iternal/model"
	"database/sql"
	log "github.com/sirupsen/logrus"
)

type CalendarRepository struct {
	storage *Storage
}

//CreateWeek sending all days of the week with a description to the database
func (r *CalendarRepository) CreateWeek(calendar map[string]string) error {
	day := [7]string{"Пн", "Вт", "Ср", "Чт", "Пт", "Сб", "Вс"}

	row := r.storage.db.QueryRow("SELECT * FROM training_calendar WHERE day = 'Сб'")
	a := &model.Calendar{}
	err := row.Scan(&a.Day, &a.Description)
	if err != sql.ErrNoRows {
		return model.ErrWeekAlreadyCreated
	}

	_, err = r.storage.db.Query("INSERT INTO training_calendar (day, description) VALUES ($1, $2), ($3, $4), ($5, $6), ($7, $8), ($9, $10), ($11, $12), ($13, $14)",
		day[0], calendar[day[0]],
		day[1], calendar[day[1]],
		day[2], calendar[day[2]],
		day[3], calendar[day[3]],
		day[4], calendar[day[4]],
		day[5], calendar[day[5]],
		day[6], calendar[day[6]],
	)
	return err
}

//GetWeek getting all days of the week with a description from the database
func (r *CalendarRepository) GetWeek() (calendar []model.Calendar, err error) {
	rows, err := r.storage.db.Query("SELECT day, description FROM training_calendar ORDER BY day")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := model.Calendar{}
		err := rows.Scan(&c.Day, &c.Description)
		if err != nil {
			log.Error(err)
			continue
		}
		calendar = append(calendar, c)
	}
	return calendar, err
}

//ChangeDay changing the description of the day of the week
func (r *CalendarRepository) ChangeDay(calendar *model.Calendar) error {
	result, err := r.storage.db.Exec("UPDATE training_calendar SET day = $1, description = $2 WHERE day = $3",
		calendar.Day,
		calendar.Description,
		calendar.Day,
	)
	if err != nil {
		return err
	}
	if count, _ := result.RowsAffected(); count != 1 {
		return model.ErrWeekNotCreated
	}
	return err
}
