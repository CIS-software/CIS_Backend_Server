package apiserver

import (
	"CIS_Backend_Server/iternal/app/apiserver/entities/handlers"
	"CIS_Backend_Server/iternal/app/apiserver/usecase/service"
	"CIS_Backend_Server/iternal/app/storage/dbstorage"
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start(cfg *Config, logger *logrus.Logger, router *mux.Router) error {
	db, err := newDB(cfg.DatabaseURL)
	if err != nil {
		return err
	}

	minioClient, err := minio.New(cfg.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: cfg.UseSSL,
	})
	if err != nil {
		logger.Fatalln(err)
	}

	//logger.Printf("%#v\n", minioClient)
	logger.Info("Client minIO is running...")
	logger.Info("The server is running...")

	defer db.Close()

	storage := dbstorage.New(db, minioClient, cfg.JwtSecretKey, cfg.AccessTokenLifetime, cfg.RefreshTokenLifetime)
	handler := handlers.New(service.New(storage))
	srv := newServer(logger, router, handler, cfg.JwtSecretKey)
	return http.ListenAndServe(cfg.BindAddr, srv)
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
