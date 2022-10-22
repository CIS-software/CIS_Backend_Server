package main

import (
	"CIS_Backend_Server/config"
	"CIS_Backend_Server/iternal/apiserver"
	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	cfg := config.New()
	logger := log.New()
	router := mux.NewRouter()

	log.Info("Reading config")
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		log.Panic(err)
	}

	if err := apiserver.Start(cfg, logger, router); err != nil {
		log.Fatal(err)
	}
}
