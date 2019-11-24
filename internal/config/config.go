package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Name string
	Db   struct {
		Host     string
		Port     string
		User     string
		Password string
	}
	App struct {
		Port     string
		LogLevel string
	}
}

func Get(path string) AppConfig {
	var config AppConfig
	configBlob, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = toml.Unmarshal(configBlob, &config); err != nil {
		log.Fatal(err)
	}
	return config
}
