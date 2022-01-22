package storage

type Storage interface {
	News() NewsRepository
}
