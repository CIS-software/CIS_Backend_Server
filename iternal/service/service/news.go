package service

import (
	"CIS_Backend_Server/iternal/model"
	"context"
	"github.com/go-playground/validator/v10"
	"strings"
)

type NewsService struct {
	service *Service
}

func (s *NewsService) Create(ctx context.Context, n *model.News) error {
	//check for file type png or jpeg
	switch n.NameSlice[len(n.NameSlice)-1] {
	case "png", "PNG":
		n.ContentType = "image/png"
	case "jpg", "jpeg", "JPG", "JPEG", "jpe", "JPE":
		n.ContentType = "image/jpeg"
	default:
		return model.ErrWrongContentType
	}

	//photo title length check
	//maximum number of characters in the photo title field in the database is 80:
	//37 - uuid and dot, 4 - extension, 39 - photo title
	var nameLength int
	for index := range n.NameSlice {
		//checking for the last element of the slice, photo extension, it is not considered
		if index == len(n.NameSlice)-1 {
			break
		}

		//number of cell characters + dot character
		nameLength = nameLength + len(n.NameSlice[index]) + 1
	}

	if nameLength > 39 {
		return model.ErrLongFileName
	}

	//title and description validation
	v := validator.New()
	if err := v.Struct(n); err != nil {
		return model.ErrTitleDescriptionNotValid
	}

	return s.service.storage.News().Create(ctx, n)
}

func (s *NewsService) Get(ctx context.Context) ([]model.News, error) {
	news, err := s.service.storage.News().Get(ctx)

	//excluding uuid from photo title
	for i := range news {
		var name string
		nameSlice := strings.Split(news[i].Name, ".")
		for index := range nameSlice {
			if index == len(nameSlice)-1 {
				name = name + nameSlice[index]
				break
			}
			if nameSlice[len(nameSlice)-2] == nameSlice[index] {
				continue
			}
			name = name + nameSlice[index] + "."
		}
		news[i].Name = name
	}

	return news, err
}

func (s *NewsService) Change(ctx context.Context, n *model.News) error {
	return s.service.storage.News().Change(ctx, n)
}

func (s *NewsService) Delete(ctx context.Context, id int) error {
	return s.service.storage.News().Delete(ctx, id)
}
