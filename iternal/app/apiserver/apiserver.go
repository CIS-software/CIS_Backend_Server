package apiserver

import (
	"CIS_Backend_Server/iternal/app/storage/dbstorage"
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start(config *Config, logger *logrus.Logger) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	logger.Info("The server is running...")

	defer db.Close()
	storage := dbstorage.New(db)
	srv := newServer(storage, logger)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
