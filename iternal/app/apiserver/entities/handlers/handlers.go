package handlers

import (
	"CIS_Backend_Server/iternal/app/apiserver/entities"
	"CIS_Backend_Server/iternal/app/apiserver/usecase"
)

type Handlers struct {
	service         usecase.Service
	handlerNews     *HandlerNews
	handlerUsers    *HandlerUser
	handlerCalendar *HandlerCalendar
}

func New(service usecase.Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) News() entities.HandlerNews {
	if h.handlerNews != nil {
		return h.handlerNews
	}

	h.handlerNews = &HandlerNews{
		handler: h,
	}
	return h.handlerNews
}

func (h *Handlers) Users() entities.HandlerUsers {
	if h.handlerUsers != nil {
		return h.handlerUsers
	}

	h.handlerUsers = &HandlerUser{
		handler: h,
	}
	return h.handlerUsers
}

func (h *Handlers) Calendar() entities.HandlersCalendar {
	if h.handlerCalendar != nil {
		return h.handlerCalendar
	}

	h.handlerCalendar = &HandlerCalendar{
		handler: h,
	}
	return h.handlerCalendar
}
