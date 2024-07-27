package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"

	"github.com/marcelofabianov/identity-gateway/config"
	"github.com/marcelofabianov/identity-gateway/pkg/postgres"
)

func TestConnect(t *testing.T) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "username",
			"POSTGRES_PASSWORD": "password",
			"POSTGRES_DB":       "dbname",
		},
		WaitingFor: wait.ForLog("listening on Unix socket").WithStartupTimeout(60 * time.Second),
	}

	c, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	require.NoError(t, err)
	defer func() {
		if err := c.Terminate(ctx); err != nil {
			t.Logf("Error terminating PostgreSQL container: %v", err)
		}
	}()

	host, err := c.Host(ctx)
	require.NoError(t, err)

	port, err := c.MappedPort(ctx, "5432")
	require.NoError(t, err)

	cfg := config.DatabaseConfig{
		Host:     host,
		Port:     port.Port(),
		User:     "username",
		Password: "password",
		Database: "dbname",
		SslMode:  "disable",
	}

	var pg *postgres.Postgres
	for i := 0; i < 3; i++ {
		pg, err = postgres.Connect(ctx, cfg)
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
	require.NoError(t, err)
	require.NotNil(t, pg)
	defer func() {
		err := pg.Close(ctx)
		require.NoError(t, err)
	}()

	err = pg.Ping(ctx)
	require.NoError(t, err)
}
