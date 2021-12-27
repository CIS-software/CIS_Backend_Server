package storage

type Storage interface {
	Events() EventsRepository
}