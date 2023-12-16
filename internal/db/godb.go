package db

import (
	"context"
	"fmt"
	"os"
	"twit-hub111/internal/config"
)

// sudo docker-compose up -d
// sudo docker-compose exec pgdb psql -U postgres -c 'CREATE DATABASE twit_hub'

func Test() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Config err: %v\n", err)
		os.Exit(1)
	}

	poolConfig, err := config.NewPoolConfig(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pool config error: %v\n", err)
		os.Exit(1)
	}

	poolConfig.MaxConns = 5
	fmt.Println(1)

	conn, err := config.NewConnection(poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect to db failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK")

	_, err = conn.Exec(context.Background(), ";")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Ping OK!")
	for i := 0; i < 5; i++ {
		go func(count int) {
			_, err = conn.Exec(context.Background(), ";")
			if err != nil {
				fmt.Fprintf(os.Stderr, "ping failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(count, "Query OK!")
			fmt.Printf("connections - MAX: %d, "+
				"Iddle: %d, "+
				"Total: %d \n",
				conn.Stat().MaxConns(),
				conn.Stat().IdleConns(),
				conn.Stat().TotalConns())
		}(i)
	}
	//conn.Close()
	select {}
}
