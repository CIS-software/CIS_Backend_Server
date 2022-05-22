package main

import (
	"CIS_Backend_Server/config"
	"CIS_Backend_Server/iternal/app/apiserver"
	"github.com/gorilla/mux"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.NewConfig()
	logger := logrus.New()

	router := mux.NewRouter()
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		logrus.Panic(err)
	}
	if err := apiserver.Start(cfg, logger, router); err != nil {
		logrus.Fatal(err)
	}
}
