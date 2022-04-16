package dbstorage

import (
	"CIS_Backend_Server/iternal/app/storage"
	"database/sql"
	_ "github.com/lib/pq" // ...
	"github.com/minio/minio-go/v7"
)

type Storage struct {
	db                 *sql.DB
	newsRepository     *NewsRepository
	usersRepository    *UsersRepository
	calendarRepository *CalendarRepository
	minioClient        *minio.Client
	SecretKey          string
	AccessLifetime     int
	RefreshLifetime    int
}

func New(db *sql.DB, minioClient *minio.Client, secretKey string, accessLifetime, refreshLifetime int) *Storage {
	return &Storage{
		db:              db,
		minioClient:     minioClient,
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

func (s *Storage) Calendar() storage.CalendarRepository {
	if s.calendarRepository != nil {
		return s.calendarRepository
	}

	s.calendarRepository = &CalendarRepository{

		storage: s,
	}
	return s.calendarRepository
}
