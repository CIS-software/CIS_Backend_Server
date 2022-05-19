package model

import (
	"io"
)

type News struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo
	TimeDate string `json:"time-date"`
}

type Photo struct {
	Payload io.Reader `json:"-"`
	Name    string    `json:"name"`
	Size    int64     `json:"-"`
	URL     string    `json:"url,omitempty"`
}
