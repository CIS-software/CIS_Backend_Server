package apiserver

import (
	"CIS_Backend_Server/config"
	"CIS_Backend_Server/iternal/handlers/handlers"
	"CIS_Backend_Server/iternal/service/service"
	"CIS_Backend_Server/iternal/storage/dbstorage"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Start(cfg *config.Config, log *log.Logger, router *mux.Router) error {
	log.Info("Database connection check")
	db, err := newDB(cfg.Postgres)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Info("Instantiating minio client")
	mc, err := minio.New(cfg.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		return err
	}

	storage := dbstorage.New(db, mc, cfg.JWT, cfg.Minio.BucketName)
	service_ := service.New(storage)
	handlers_ := handlers.New(service_)
	srv := newServer(log, router, handlers_, cfg.JWT.SecretKey)
	log.Info("Server start...")
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
