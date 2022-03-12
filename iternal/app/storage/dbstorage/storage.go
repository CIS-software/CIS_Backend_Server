package dbstorage

import (
	"CIS_Backend_Server/iternal/app/storage"
	"database/sql"
	_ "github.com/lib/pq" // ...
)

type Storage struct {
	db              *sql.DB
	newsRepository  *NewsRepository
	usersRepository *UsersRepository
	SecretKey       string
	AccessLifetime  int
	RefreshLifetime int
}

func New(db *sql.DB, secretKey string, accessLifetime int, refreshLifetime int) *Storage {
	return &Storage{
		db:              db,
		SecretKey:       secretKey,
		AccessLifetime:  accessLifetime,
		RefreshLifetime: refreshLifetime,
	}
}

func (s *Storage) Users() storage.UsersRepository {
	if s.usersRepository != nil {
		return s.usersRepository
	}

	s.usersRepository = &UsersRepository{

		storage: s,
	}
	return s.usersRepository
}

func (s *Storage) News() storage.NewsRepository {
	if s.newsRepository != nil {
		return s.newsRepository
	}

	s.newsRepository = &NewsRepository{

		storage: s,
	}
	return s.newsRepository
}
