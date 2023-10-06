package config

import (
	"github.com/joho/godotenv"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database   Database
	ServerPort int `envconfig:"SERVER_PORT" default:"80"`
}

type Database struct {
	Host         string `envconfig:"DATABASE_HOST" required:"true"`
	Port         int    `envconfig:"DATABASE_PORT" required:"true"`
	User         string `envconfig:"DATABASE_USER" required:"true"`
	Password     string `envconfig:"DATABASE_PASSWORD" required:"true"`
	DatabaseName string `envconfig:"DATABASE_NAME" required:"true"`
}

func NewParsedConfig() (Config, error) {
	err := godotenv.Load(".env")
	cnf := Config{}
	if err != nil {
		return cnf, err
	}
	err = envconfig.Process("", &cnf)
	return cnf, err
}
