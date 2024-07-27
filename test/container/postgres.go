package container

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
)

type PostgresContainer struct {
	Container *postgres.PostgresContainer
	ConnStr   string
}

func NewPostgresContainer(t *testing.T, ctx context.Context) *PostgresContainer {
	container, err := postgres.Run(
		ctx,
		"postgres:16",
		postgres.WithDatabase("dbname"),
		postgres.WithUsername("username"),
		postgres.WithPassword("password"),
	)
	if err != nil {
		t.Fatalf("01. Error creating PostgreSQL container: %v", err)
	}

	t.Cleanup(func() {
		if container != nil {
			if err := container.Terminate(ctx); err != nil {
				t.Logf("02. Error terminating PostgreSQL container: %v", err)
			}
		}
	})

	connStr, err := container.ConnectionString(ctx, "sslmode=disable", "timezone=UTC", "application_name=test")
	if err != nil {
		t.Fatalf("03. Error getting connection string: %v", err)
	}

	var db *sql.DB
	for i := 0; i < 1; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			t.Fatalf("04. Error opening SQL connection: %v", err)
		}
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("05. Error pinging SQL database: %v", err)
	}

	return &PostgresContainer{
		Container: container,
		ConnStr:   connStr,
	}
}
