package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
)

type Config struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
	Timeout  int    `yaml:"Timeout"`
}

func ReadConfig() (*Config, error) {
	yFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		return nil, err
	}

	var cfg Config
	_ = yaml.Unmarshal(yFile, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func NewPoolConfig(cfg *Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable&connect_timeout=%d",
		"postgres",
		url.QueryEscape(cfg.Username),
		url.QueryEscape(cfg.Password),
		cfg.Host,
		cfg.Port,
		cfg.DbName,
		cfg.Timeout)

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
