package storage

import (
	"CIS_Backend_Server/iternal/model"
	"database/sql"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UsersRepository struct {
	storage *Storage
}

func (r *UsersRepository) CreateUser(a *model.UserAuth, u *model.User) error {
	var email string
	err := r.storage.db.QueryRow(
		"SELECT email FROM user_auth WHERE email = $1", a.Email).Scan(&email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	//check if email exists in database
	if a.Email == email {
		return model.ErrEmailIsBusy
	}

	//create encrypted password
	enc, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	//type byte to type string
	a.EncryptedPassword = string(enc)

	//sending email and encrypted password to database
	err = r.storage.db.QueryRow(
		"INSERT INTO user_auth (email, encrypted_password) VALUES ($1, $2) RETURNING id",
		a.Email,
		a.EncryptedPassword,
	).Scan(&a.Id)
	if err != nil {
		return err
	}

	//sending profile data to the database
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
	//getting user profile from database
	row := r.storage.db.QueryRow("SELECT * FROM user_profile WHERE user_id = $1", id)
	user = new(model.User)
	err = row.Scan(
		&user.UserId, &user.Name, &user.Surname, &user.Town,
		&user.Age, &user.Belt, &user.Weight, &user.IdIKO)

	//error handling if the user is not found by id or another error
	if err == sql.ErrNoRows {
		return nil, model.ErrUserNotFound
	} else if err != nil {
		return nil, err
	}

	return user, err
}

func (r *UsersRepository) Login(a *model.UserAuth, t *model.Tokens) error {
	//getting id and encrypted password from database
	row := r.storage.db.QueryRow("SELECT id, encrypted_password FROM user_auth WHERE email = $1",
		a.Email,
	).Scan(&a.Id, &a.EncryptedPassword)
	if row != nil {
		return model.ErrWrongEmailOrPassword
	}

	//comparing an encrypted password with an unencrypted one
	err := bcrypt.CompareHashAndPassword([]byte(a.EncryptedPassword), []byte(a.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return model.ErrWrongEmailOrPassword
	} else if err != nil {
		return err
	}

	//creation of access token
	t.AccessToken, err = createToken(a.Id, r.storage.AccessLifetime, r.storage.AccessKey)
	if err != nil {
		return err
	}

	//creation of refresh token
	t.RefreshToken, err = createToken(a.Id, r.storage.RefreshLifetime, r.storage.RefreshKey)
	if err != nil {
		return err
	}

	t.TokenId = a.Id

	return err
}

func (r *UsersRepository) UpdateTokens(t *model.Tokens) error {
	//token validation
	_, err := jwt.ParseWithClaims(t.RefreshToken, t, func(token *jwt.Token) (interface{}, error) {
		return []byte(r.storage.RefreshKey), nil
	})
	if err != nil {
		return model.ErrWrongToken
	}

	//create access token
	t.AccessToken, err = createToken(t.TokenId, r.storage.AccessLifetime, r.storage.AccessKey)
	if err != nil {
		return err
	}

	//create refresh token
	t.RefreshToken, err = createToken(t.TokenId, r.storage.RefreshLifetime, r.storage.RefreshKey)
	if err != nil {
		return err
	}

	return err
}

func createToken(id, lifetime int, secretKey string) (string, error) {
	claims := &model.Tokens{
		TokenId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetime)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))

	return tokenString, err
}
