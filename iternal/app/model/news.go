package model

type News struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
}
