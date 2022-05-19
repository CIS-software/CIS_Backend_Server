package main

import (
	"CIS_Backend_Server/iternal/app/apiserver"
	"github.com/BurntSushi/toml"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

var (
	configPath string = "go/bin/config.toml"
)

func main() {
	config := apiserver.NewConfig()
	logger := logrus.New()
	router := mux.NewRouter()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Panic(err)
	}
	if err := apiserver.Start(config, logger, router); err != nil {
		logrus.Fatal(err)
	}
}
