package handlers

import (
	"CIS_Backend_Server/iternal/handlers"
	"CIS_Backend_Server/iternal/service"
)

type Handlers struct {
	service         service.Service
	handlerNews     *HandlerNews
	handlerUsers    *HandlerUser
	handlerCalendar *HandlerCalendar
}

func New(service service.Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) News() handlers.HandlerNews {
	if h.handlerNews != nil {
		return h.handlerNews
	}

	h.handlerNews = &HandlerNews{
		handler: h,
	}
	return h.handlerNews
}

func (h *Handlers) Users() handlers.HandlerUsers {
	if h.handlerUsers != nil {
		return h.handlerUsers
	}

	h.handlerUsers = &HandlerUser{
		handler: h,
	}
	return h.handlerUsers
}

func (h *Handlers) Calendar() handlers.HandlerCalendar {
	if h.handlerCalendar != nil {
		return h.handlerCalendar
	}

	h.handlerCalendar = &HandlerCalendar{
		handler: h,
	}
	return h.handlerCalendar
}
