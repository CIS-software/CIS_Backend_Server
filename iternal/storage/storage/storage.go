package storage

import (
	"CIS_Backend_Server/config"
	"CIS_Backend_Server/iternal/storage"
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
	bucketName         string
	ConfigJWT
}

type ConfigJWT struct {
	AccessKey       string
	RefreshKey      string
	AccessLifetime  int
	RefreshLifetime int
}

func New(db *sql.DB, mc *minio.Client, cfg config.JWT, bucketName string) *Storage {
	return &Storage{
		db:          db,
		minioClient: mc,
		bucketName:  bucketName,
		ConfigJWT: ConfigJWT{
			AccessKey:       cfg.SecretKeyAccess,
			RefreshKey:      cfg.SecretKeyRefresh,
			AccessLifetime:  cfg.AccessTokenLifetime,
			RefreshLifetime: cfg.RefreshTokenLifetime,
		},
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
