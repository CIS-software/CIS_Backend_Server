package entities

import (
	"net/http"
)

type Handlers interface {
	News() HandlerNews
	Users() HandlerUsers
	Calendar() HandlersCalendar
}

type HandlerNews interface {
	HandleCreateNews() http.HandlerFunc
	HandleGetNews() http.HandlerFunc
	HandleUpdateNews() http.HandlerFunc
	HandleDeleteNews() http.HandlerFunc
}

type HandlerUsers interface {
	HandleGetUser() http.HandlerFunc
	HandleCreateUser() http.HandlerFunc
	HandleLogin() http.HandlerFunc
	HandleUpdateTokens() http.HandlerFunc
}

type HandlersCalendar interface {
	HandleCreateTrainingWeek() http.HandlerFunc
	HandleGetTrainingCalendar() http.HandlerFunc
	HandleUpdateTraining() http.HandlerFunc
}
