package dbstorage

import (
	"CIS_Backend_Server/iternal/app/model"
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository struct {
	storage *Storage
}

func (r *UsersRepository) CreateUser(u *model.User) error {
	return r.storage.db.QueryRow(
		"INSERT INTO user_profile (user_id, name, surname, patronymic, town, age, weight, belt, id_iko) "+
			"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		u.UserID,
		u.Name,
		u.Surname,
		u.Patronymic,
		u.Town,
		u.Age,
		u.Weight,
		u.Belt,
		u.IdIKO,
	).Err()
}

func (r *UsersRepository) GetUser(id int) (user *model.User, err error) {
	row := r.storage.db.QueryRow("SELECT * FROM user_profile WHERE user_id = $1", id)
	user = new(model.User)
	err = row.Scan(&user.UserID, &user.Name, &user.Surname, &user.Patronymic, &user.Town, &user.Age, &user.Belt, &user.Weight, &user.IdIKO)

	if err == sql.ErrNoRows {
		err = errors.New("user not found")
		return user, err
	} else if err != nil {
		logrus.Panic(err)
	}

	return user, err
}

func (r *UsersRepository) CreateUserAuth(u *model.UserAuth) error {
	enc, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		logrus.Panic(err)
		return err
	}
	u.EncryptedPassword = string(enc)

	row := r.storage.db.QueryRow(
		"INSERT INTO user_auth (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		u.Email,
		u.EncryptedPassword,
	).Scan(&u.Id)
	return row
}

func (r *UsersRepository) Login(u *model.UserAuth) error {
	row := r.storage.db.QueryRow("SELECT id, encrypted_password FROM user_auth WHERE email = $1",
		u.Email,
	).Scan(&u.Id, &u.EncryptedPassword)

	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(u.Password))
	if row != nil || err != nil {
		return errors.New("wrong email or password")
	}

	u.AccessToken, err = model.CreateToken(u.Id, r.storage.AccessLifetime, u.Email, r.storage.SecretKey)
	if err != nil {
		return err
	}
	u.RefreshToken, err = model.CreateToken(u.Id, r.storage.RefreshLifetime, u.Email, r.storage.SecretKey)
	if err != nil {
		return err
	}

	_, err = r.storage.db.Exec(
		"UPDATE user_auth SET access_token = $1, refresh_token = $2 WHERE id = $3",
		u.AccessToken,
		u.RefreshToken,
		u.Id,
	)
	return err
}

func (r *UsersRepository) UpdateTokens(u *model.UserAuth) error {
	_, err := jwt.ParseWithClaims(u.RefreshToken, u, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.storage.SecretKey), nil
	})
	if err != nil {
		return errors.New("wrong token")
	}

	u.AccessToken, err = model.CreateToken(u.Id, r.storage.AccessLifetime, u.Email, r.storage.SecretKey)
	if err != nil {
		return err
	}
	u.RefreshToken, err = model.CreateToken(u.Id, r.storage.RefreshLifetime, u.Email, r.storage.SecretKey)
	if err != nil {
		return err
	}

	_, err = r.storage.db.Exec(
		"UPDATE user_auth SET access_token = $1, refresh_token = $2 WHERE id = $3",
		u.AccessToken,
		u.RefreshToken,
		u.Id,
	)
	return err
}
