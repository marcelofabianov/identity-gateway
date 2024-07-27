package container

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go/modules/postgres"

	"github.com/marcelofabianov/identity-gateway/config"
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
		t.Fatalf("Error creating PostgreSQL container: %v", err)
	}

	t.Cleanup(func() {
		if container != nil {
			if err := container.Terminate(ctx); err != nil {
				t.Logf("Error terminating PostgreSQL container: %v", err)
			}
		}
	})

	connStr, err := container.ConnectionString(ctx, "sslmode=disable", "timezone=UTC", "application_name=test")
	if err != nil {
		t.Fatalf("Error getting connection string: %v", err)
	}

	var db *sql.DB
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			t.Fatalf("Error opening SQL connection: %v", err)
		}
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(2 * time.Second)
	}

	if err := db.Ping(); err != nil {
		t.Fatalf("Error pinging SQL database: %v", err)
	}

	scriptPath := config.GetCurrentPathRelative("schema.sql")
	if err := executeSQLScript(db, scriptPath); err != nil {
		t.Fatalf("Error executing SQL script: %v", err)
	}

	return &PostgresContainer{
		Container: container,
		ConnStr:   connStr,
	}
}

func executeSQLScript(db *sql.DB, filename string) error {
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return fmt.Errorf("error determining absolute path: %w", err)
	}

	fmt.Printf("Opening SQL file: %s\n", absPath)

	file, err := os.Open(absPath)
	if err != nil {
		return fmt.Errorf("error opening SQL file: %w", err)
	}
	defer file.Close()

	sqlBytes, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading SQL file: %w", err)
	}
	sqlScript := string(sqlBytes)

	fmt.Println("Executing SQL script")
	_, err = db.Exec(sqlScript)
	if err != nil {
		return fmt.Errorf("error executing SQL script: %w", err)
	}

	return nil
}
