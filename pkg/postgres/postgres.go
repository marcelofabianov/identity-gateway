package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/marcelofabianov/identity-gateway/config"
)

type Postgres struct {
	conn *sql.DB
}

func NewPostgres(conn *sql.DB) *Postgres {
	return &Postgres{conn: conn}
}

func (p *Postgres) Conn() *sql.DB {
	return p.conn
}

func (p *Postgres) Close(ctx context.Context) error {
	return p.conn.Close()
}

func (p *Postgres) Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return p.conn.PingContext(ctx)
}

func Connect(ctx context.Context, cfg config.DatabaseConfig) (*Postgres, error) {
	dsn := FormatDSN(cfg)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := conn.PingContext(ctx); err != nil {
		return nil, err
	}

	return NewPostgres(conn), nil
}

func FormatDSN(cfg config.DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SslMode)
}
