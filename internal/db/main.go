package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"net/url"
	"os"
)

func main() {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s?sslmode=disable&connect_timeout=%d//",
		"postgres",
		url.QueryEscape("postgres"),
		url.QueryEscape("123"),
		"localhost",
		"54320",
		"twit-hub",
		5)

	ctx, _ := context.WithCancel(context.Background())

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect to db failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK")

	//сделаем пустой запрос
	err = conn.Ping(ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ping failed: %v\n", err)
		os.Exit(1)
	}
}
