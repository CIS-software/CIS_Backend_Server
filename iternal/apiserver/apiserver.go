package apiserver

import (
	"CIS_Backend_Server/config"
	"CIS_Backend_Server/iternal/handlers/handlers"
	"CIS_Backend_Server/iternal/service/service"
	"CIS_Backend_Server/iternal/storage/storage"
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
	//open db and check connection
	db, err := newDB(cfg.Postgres)
	if err != nil {
		return err
	}
	defer db.Close()

	log.Info("Instantiating minio client")
	//instantiate minio client with options
	mc, err := minio.New(cfg.Minio.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.Minio.AccessKeyID, cfg.Minio.SecretAccessKey, ""),
		Secure: cfg.Minio.UseSSL,
	})
	if err != nil {
		return err
	}

	//instantiate storage, service, handlers, server
	storageNew := storage.New(db, mc, cfg.JWT, cfg.BucketName)
	serviceNew := service.New(storageNew)
	handlersNew := handlers.New(serviceNew)
	srv := serverNew(log, router, handlersNew, cfg.SecretKeyAccess)

	log.Info("Server start on port: ", cfg.BindAddr)
	return http.ListenAndServe(cfg.BindAddr, srv)
}

// newDB opening a database according to the specified parameters, checking the connection
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
