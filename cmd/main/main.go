package main

import (
	"CIS_Backend_Server/iternal/app/apiserver"
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	logger := logrus.New()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		logrus.Fatal(err)
	}

	if err := apiserver.Start(config, logger); err != nil {
		logrus.Fatal(err)
	}
}
