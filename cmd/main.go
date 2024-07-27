package main

import (
	"context"
	"log"

	"github.com/marcelofabianov/identity-gateway/config"
	"github.com/marcelofabianov/identity-gateway/pkg/postgres"
	"github.com/marcelofabianov/identity-gateway/pkg/zap"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	logger, err := zap.NewLogger(cfg.Log)
	if err != nil {
		log.Fatalf("error creating logger: %v", err)
	}
	defer logger.Close()

	ctx := context.Background()

	db, err := postgres.Connect(ctx, cfg.Db)
	if err != nil {
		log.Fatal("error connecting to database")
	}
	defer func() {
		if err := db.Close(ctx); err != nil {
			log.Fatal("error closing database connection")
		}
	}()

	logger.Info("starting application")
}
