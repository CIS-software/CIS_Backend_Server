package service

import (
	"CIS_Backend_Server/iternal/model"
	"context"
	"github.com/go-playground/validator/v10"
)

type NewsService struct {
	service *Service
}

func (s *NewsService) Create(ctx context.Context, n *model.News) error {
	//validation of received data from the client
	if err := newsValidation(n); err != nil {
		return err
	}

	return s.service.storage.News().Create(ctx, n)
}

func (s *NewsService) Get(ctx context.Context, id int) ([]model.News, error) {
	news, err := s.service.storage.News().Get(ctx, id)

	//excluding uuid from photo title
	for i := range news {
		var name = news[i].NameWithUUID
		if string(name[len(name)-5]) == "." {
			name = name[:len(name)-4-37] + name[len(name)-4:]
		} else {
			name = name[:len(name)-3-37] + name[len(name)-3:]
		}
		news[i].Name = name
	}

	return news, err
}

func (s *NewsService) Change(ctx context.Context, n *model.News) error {
	//validation of received data from the client
	if err := newsValidation(n); err != nil {
		return err
	}

	return s.service.storage.News().Change(ctx, n)
}

func (s *NewsService) Delete(ctx context.Context, id int) error {
	return s.service.storage.News().Delete(ctx, id)
}

//newsValidation validation of file extension, photo name length, news title and description
func newsValidation(n *model.News) error {
	//check for file type png or jpeg
	switch n.Name[len(n.Name)-4:] {
	case ".png", ".PNG":
		n.ContentType = "image/png"
	case ".jpg", "jpeg", ".JPG", "JPEG", ".jpe", ".JPE":
		if n.Name[len(n.Name)-5:] != ".jpeg" && n.Name[len(n.Name)-5:] != ".JPEG" {
			return model.ErrWrongContentType
		}
		n.ContentType = "image/jpeg"
	default:
		return model.ErrWrongContentType
	}

	//title and description validation
	v := validator.New()
	if err := v.Struct(n); err != nil {
		return model.ErrNewsNotValid
	}

	return nil
}
