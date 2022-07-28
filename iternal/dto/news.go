package dto

type News struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	TimeDate    string `json:"time-date"`
}
