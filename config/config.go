package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP HTTP
		Log  Log
	}

	HTTP struct {
		Port string `env:"HTTP_PORT" env-default:"8080"`
	}

	Log struct {
		Level string `env:"LOG_LEVEL" env-default:"info"`
	}
)

func Get() *Config {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		log.Fatal("failed to read env", err)
	}

	return &config
}
