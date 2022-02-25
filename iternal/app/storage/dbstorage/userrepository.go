package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"github.com/sirupsen/logrus"
)

type UsersRepository struct {
	storage *Storage
}

func (r *UsersRepository) CreateUser(u *model.User) error {
	return r.storage.db.QueryRow(
		"INSERT INTO users (name, surname, patronymic, town, age, weight) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, belt, id_iko",
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Town,
		u.Age,
		u.Weight,
	).Scan(&u.Id, &u.Belt, &u.IdIKO)
}

func (r *UsersRepository) GetUsers() (users []model.User, err error) {
	rows, err := r.storage.db.Query("SELECT * FROM users")
	if err != nil {
		logrus.Panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		u := model.User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Surname, &u.Patronymic, &u.Town, &u.Age, &u.Belt, &u.Weight, &u.IdIKO)
		if err != nil {
			logrus.Info(err)
			continue
		}
		users = append(users, u)
	}
	return users, err
}

func (r *UsersRepository) Login(u *model.User) (uint64, error) {
	return 0, nil
}
