package domens

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v2"
	"net/url"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Timeout  int
}

func (cfg *Config) ReadConfig() error {

	// Считываем содержимое файла
	data, err := os.ReadFile("conf.yaml")
	if err != nil {
		return err
	}

	// Производим анмаршалинг данных в структуру Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	return nil
}

func NewPoolConfig(cfg *Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d//",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		5)

	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		return nil, err
	}

	return poolConfig, nil
}

func NewConnection(poolConfig *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
