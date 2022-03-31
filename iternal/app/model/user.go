package model

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type User struct {
	UserId  int     `json:"userId"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Town    string  `json:"town"`
	Age     string  `json:"age"`
	Belt    string  `json:"belt"`
	Weight  float32 `json:"weight"`
	IdIKO   string  `json:"id_iko"`
}

type UserAuth struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Password          string `json:"-"`
	EncryptedPassword string `json:"-"`
}

type Tokens struct {
	TokenId      int    `json:"id,omitempty"`
	AccessToken  string `json:"access-token,omitempty"`
	RefreshToken string `json:"refresh-token,omitempty"`
	jwt.StandardClaims
}

type UserData struct {
	Id      int     `json:"id"`
	Email   string  `json:"email"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Town    string  `json:"town"`
	Age     string  `json:"age"`
	Belt    string  `json:"belt"`
	Weight  float32 `json:"weight"`
	IdIKO   string  `json:"id_iko"`
}

func CreateToken(id, lifetime int, secretKey string) (string, error) {
	claims := &Tokens{
		TokenId: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetime)).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, err
}
