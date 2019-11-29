package config

import (
	"github.com/BurntSushi/toml"
	"github.com/lekhacman/dms/internal/store"
	"io/ioutil"
	"log"
)

type AppConfig struct {
	Name string
	Db   store.DbSpec
	App  struct {
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
