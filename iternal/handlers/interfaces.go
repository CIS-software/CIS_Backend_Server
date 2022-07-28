package handlers

import (
	"net/http"
)

type Handlers interface {
	News() HandlerNews
	Users() HandlerUsers
	Calendar() HandlerCalendar
}

type HandlerNews interface {
	Create() http.HandlerFunc
	Get() http.HandlerFunc
	Change() http.HandlerFunc
	Delete() http.HandlerFunc
}

type HandlerUsers interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Login() http.HandlerFunc
	RefreshTokens() http.HandlerFunc
}

type HandlerCalendar interface {
	CreateWeek() http.HandlerFunc
	GetWeek() http.HandlerFunc
	ChangeDay() http.HandlerFunc
}
