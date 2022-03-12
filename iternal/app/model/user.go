package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type User struct {
	UserID     int     `json:"user-id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic string  `json:"patronymic"`
	Town       string  `json:"town"`
	Age        int     `json:"age"`
	Belt       string  `json:"belt"`
	Weight     float32 `json:"weight"`
	IdIKO      string  `json:"id_iko"`
}

type UserAuth struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"-"`
	EncryptedPassword string `json:"-"`
	AccessToken       string `json:"access-token,omitempty"`
	RefreshToken      string `json:"refresh-token,omitempty"`
	jwt.StandardClaims
}

type Tokens struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

func CreateToken(id, lifetime int, email, secretKey string) (string, error) {
	claims := &UserAuth{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			Subject:   string(id),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetime)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
