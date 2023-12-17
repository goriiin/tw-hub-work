package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
	"twit-hub111/internal/config"
)

type Storage struct {
	db *pgxpool.Pool
}

func New() (*Storage, error) {
	const op = "storage.postgres.New"

	cfg, err := config.ReadConfig()
	if err != nil {
		_ = fmt.Errorf("%s - config err: %w", op, err)
		os.Exit(1)
	}

	poolConfig, err := config.NewPoolConfig(cfg)
	if err != nil {
		_ = fmt.Errorf("%s - Pool config error: %w", op, err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 5

	conn, err := config.NewConnection(poolConfig)
	if err != nil {
		_ = fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{conn}, nil
}
