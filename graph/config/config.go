package config

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	MongoDBEndpoint  string `env:"MONGODB_ENDPOINT" envDefault:"mongodb://localhost:27017"`
	MongoDBName      string `env:"MONGODB_NAME"`
	MongoDBTableName string `env:"MONGODB_TABLE_NAME"`
	PORT             string `env:"PORT"`
	HOST             string `env:"HOST"`
	PATH             string `env:"PATH_URL"`
}

func Get() *Config {
	appConfig := &Config{}
	if err := env.Parse(appConfig); err != nil {
		panic(err)
	}

	return appConfig
}
