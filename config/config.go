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
	Api      ApiServerConfig
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

type ApiServerConfig struct {
	Host     string
	Port     string
	CertPath string
	KeyPath  string
	TimeOut  string
	Healthz  string
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
		Api: ApiServerConfig{
			Host:     os.Getenv("IG_API_HOST"),
			Port:     os.Getenv("IG_API_PORT"),
			CertPath: os.Getenv("IG_API_CERT_PATH"),
			KeyPath:  os.Getenv("IG_API_KEY_PATH"),
			TimeOut:  os.Getenv("IG_API_TIMEOUT"),
			Healthz:  os.Getenv("IG_API_HEALTHZ"),
		},
	}

	return cfg, nil
}
