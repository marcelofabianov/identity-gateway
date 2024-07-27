package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Timezone string
	Name     string
	Db       DatabaseConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SslMode  string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Env:      os.Getenv("ENV"),
		Timezone: os.Getenv("TZ"),
		Name:     os.Getenv("NAME"),
		Db: DatabaseConfig{
			Host:     os.Getenv("PG_HOST"),
			Port:     os.Getenv("PG_PORT"),
			User:     os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Database: os.Getenv("PG_DATABASE"),
			SslMode:  os.Getenv("PG_SSL_MODE"),
		},
	}

	return cfg, nil
}
