package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env      string
	Timezone string
	Name     string
	Log      LogConfig
	Db       DatabaseConfig
}

type LogConfig struct {
	Level    string
	Format   string
	Output   string
	FilePath string
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
		Log: LogConfig{
			Level:    os.Getenv("IG_LOG_LEVEL"),
			Format:   os.Getenv("IG_LOG_FORMAT"),
			Output:   os.Getenv("IG_LOG_OUTPUT"),
			FilePath: os.Getenv("IG_LOG_PATH"),
		},
		Db: DatabaseConfig{
			Host:     os.Getenv("IG_DB_HOST"),
			Port:     os.Getenv("IG_DB_PORT"),
			User:     os.Getenv("IG_DB_USER"),
			Password: os.Getenv("IG_DB_PASSWORD"),
			Database: os.Getenv("IG_DB_DATABASE"),
			SslMode:  os.Getenv("IG_DB_MODE"),
		},
	}

	return cfg, nil
}
