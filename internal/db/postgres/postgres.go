package postgres

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	Username string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASS" default:"postgres"`
	Database string `envconfig:"DB_DATABASE" default:"postgres"`
	Timeout  int    `envconfig:"DB_TIMEOUT" default:"5"`
}

func NewPool(cfg *Config) (*pgxpool.Pool, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Timeout,
	)

	poolCfg, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), poolCfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}

	if err := connPool.Ping(context.Background()); err != nil {
		connPool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return connPool, nil
}
