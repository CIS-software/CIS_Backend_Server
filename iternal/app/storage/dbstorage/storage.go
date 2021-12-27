package dbstorage

import (
	"CIS_Backend_Server/iternal/app/storage"
	"database/sql"
	_ "github.com/lib/pq" // ...
)

type Storage struct {
	db 					*sql.DB
	eventsRepository 	*EventsRepository
}

func New(db *sql.DB) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Events() storage.EventsRepository {
	if s.eventsRepository != nil {
		return s.eventsRepository
	}

	s.eventsRepository = &EventsRepository{

		storage: s,
	}
	return s.eventsRepository
}