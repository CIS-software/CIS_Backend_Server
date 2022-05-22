package apiserver

import (
	"CIS_Backend_Server/config"
	"CIS_Backend_Server/iternal/app/apiserver/entities/handlers"
	"CIS_Backend_Server/iternal/app/apiserver/usecase/service"
	"CIS_Backend_Server/iternal/app/storage/dbstorage"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start(cfg *config.Config, logger *logrus.Logger, router *mux.Router) error {
	db, err := newDB(cfg.Postgres)
	if err != nil {
		return err
	}

	mc, err := minio.New(cfg.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		logger.Fatalln(err)
	}
	defer db.Close()

	storage := dbstorage.New(db, mc, cfg.JWT, cfg.Minio.BucketName)
	handler := handlers.New(service.New(storage))
	srv := newServer(logger, router, handler, cfg.JWT.SecretKey)
	logger.Info("Configuration read successfully")
	logger.Info("Server start...")
	return http.ListenAndServe(cfg.BindAddr, srv)
}

func newDB(pq config.Postgres) (*sql.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s", pq.User, pq.Password, pq.Host, pq.DBName, pq.SSL)
	db, err := sql.Open("postgres", dbURL)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
