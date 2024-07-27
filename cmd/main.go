package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/identity-gateway/config"
	"github.com/marcelofabianov/identity-gateway/pkg/logger"
	"github.com/marcelofabianov/identity-gateway/pkg/postgres"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	logger, err := logger.NewLogger(cfg.Log)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
	defer logger.Close()

	ctx := context.Background()

	db, err := postgres.Connect(ctx, cfg.Db)
	if err != nil {
		logger.Fatal("error connecting to database", logger.FieldError(err))
	}
	defer func() {
		if err := db.Close(ctx); err != nil {
			log.Fatal("error closing database connection")
		}
	}()

	logger.Info("starting application")
}
