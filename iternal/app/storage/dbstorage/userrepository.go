package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
)

type UsersRepository struct {
	storage *Storage
}

func (r *UsersRepository) CreateUser(u *model.User) error {
	return r.storage.db.QueryRow(
		"INSERT INTO users (name, surname, patronymic, town, age, weight, belt, id_iko) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Town,
		u.Age,
		u.Weight,
		u.Belt,
		u.IdIKO,
	).Scan(&u.Id)
}

func (r *UsersRepository) GetUsers(id int) (user *model.User, err error) {
	row := r.storage.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	user = new(model.User)
	err = row.Scan(&user.Id, &user.Name, &user.Surname, &user.Patronymic, &user.Town, &user.Age, &user.Belt, &user.Weight, &user.IdIKO)

	if err == sql.ErrNoRows {
		err = errors.New("User not found")
		return user, err
	} else if err != nil {
		logrus.Panic(err)
	}

	return user, err
}

func (r *UsersRepository) Login(u *model.User) (uint64, error) {
	return 0, nil
}
