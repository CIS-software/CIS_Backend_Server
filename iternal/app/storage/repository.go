package storage

import "CIS_Backend_Server/iternal/app/model"

type EventsRepository interface {
	CreateEvents(events *model.Events) error
	GetEvents() ([]model.Events, error)
}
