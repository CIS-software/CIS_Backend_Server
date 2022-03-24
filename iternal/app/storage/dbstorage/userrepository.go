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

func (r *UsersRepository) CreateUser(a *model.UserAuth, u *model.User) error {
	enc, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.EncryptedPassword = string(enc)

	err = r.storage.db.QueryRow(
		"INSERT INTO user_auth (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		a.Email,
		a.EncryptedPassword,
	).Scan(&a.Id)
	if err != nil {
		return err
	}

	return r.storage.db.QueryRow(
		"INSERT INTO user_profile (user_id, name, surname, town, age, weight, belt, id_iko) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		a.Id,
		u.Name,
		u.Surname,
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
	err = row.Scan(&user.UserId, &user.Name, &user.Surname, &user.Town, &user.Age, &user.Belt, &user.Weight, &user.IdIKO)

	if err == sql.ErrNoRows {
		err = errors.New("user not found")
		return user, err
	} else if err != nil {
		logrus.Error(err)
	}

	return user, err
}

func (r *UsersRepository) Login(a *model.UserAuth, t *model.Tokens) error {
	row := r.storage.db.QueryRow("SELECT id, encrypted_password FROM user_auth WHERE email = $1",
		a.Email,
	).Scan(&a.Id, &a.EncryptedPassword)
	if row != nil {
		return errors.New("wrong email or password")
	}

	err := bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(a.Password))
	if err != nil {
		return errors.New("wrong email or password")
	}

	t.AccessToken, err = model.CreateToken(a.Id, r.storage.AccessLifetime, r.storage.SecretKey)
	if err != nil {
		return err
	}
	t.RefreshToken, err = model.CreateToken(a.Id, r.storage.RefreshLifetime, r.storage.SecretKey)
	if err != nil {
		return err
	}

	_, err = r.storage.db.Exec(
		"UPDATE user_auth SET access_token = $1, refresh_token = $2 WHERE id = $3",
		t.AccessToken,
		t.RefreshToken,
		a.Id,
	)
	return err
}

func (r *UsersRepository) UpdateTokens(t *model.Tokens) error {
	_, err := jwt.ParseWithClaims(t.RefreshToken, t, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.storage.SecretKey), nil
	})
	if err != nil {
		return errors.New("wrong token")
	}

	t.AccessToken, err = model.CreateToken(t.TokenId, r.storage.AccessLifetime, r.storage.SecretKey)
	if err != nil {
		return err
	}
	t.RefreshToken, err = model.CreateToken(t.TokenId, r.storage.RefreshLifetime, r.storage.SecretKey)
	if err != nil {
		return err
	}
	_, err = r.storage.db.Exec(
		"UPDATE user_auth SET access_token = $1, refresh_token = $2 WHERE id = $3",
		t.AccessToken,
		t.RefreshToken,
		t.TokenId,
	)
	return err
}
