package model

import (
	"io"
)

type News struct {
	Id          int    `json:"id"`
	Title       string `json:"title" validate:"required,max=50"`
	Description string `json:"description" validate:"max=1500"`
	Photo
	TimeDate string `json:"time-date"`
}

type Photo struct {
	Payload     io.Reader `json:"-"`
	Name        string    `json:"name"`
	NameSlice   []string  `json:"name-slice"`
	ContentType string    `json:"content-type"`
	Size        int64     `json:"-"`
	URL         string    `json:"url,omitempty"`
}
