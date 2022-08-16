package model

import (
	"github.com/golang-jwt/jwt"
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
	Email             string `json:"email" validate:"email,max=50"`
	Password          string `json:"-" validate:"min=4,max=50"`
	EncryptedPassword string `json:"-"`
}

type Tokens struct {
	TokenId      int    `json:"id,omitempty"`
	AccessToken  string `json:"access-token,omitempty"`
	RefreshToken string `json:"refresh-token,omitempty"`
	jwt.StandardClaims
}
