package service

import (
	"CIS_Backend_Server/iternal/model"
	"github.com/go-playground/validator/v10"
)

type UsersService struct {
	service *Service
}

func (s *UsersService) CreateUser(userAuth *model.UserAuth, user *model.User) error {
	//email and password validation
	v := validator.New()
	if err := v.Struct(userAuth); err != nil {
		return model.ErrEmailPasswordNotValid
	}

	return s.service.storage.Users().CreateUser(userAuth, user)
}

func (s *UsersService) GetUser(id int) (users *model.User, err error) {
	return s.service.storage.Users().GetUser(id)
}

func (s *UsersService) Login(userAuth *model.UserAuth, tokens *model.Tokens) error {
	//email and password validation
	v := validator.New()
	if err := v.Struct(userAuth); err != nil {
		return model.ErrEmailPasswordNotValid
	}

	return s.service.storage.Users().Login(userAuth, tokens)
}

func (s *UsersService) UpdateTokens(tokens *model.Tokens) error {
	return s.service.storage.Users().UpdateTokens(tokens)
}
