package db

import (
	"context"
	"fmt"
	"os"
	"twit-hub111/internal/db/domens"
)

func main() {
	var cfg domens.Config
	if err := cfg.ReadConfig(); err != nil {

	}

	poolConfig, err := domens.NewPoolConfig(&cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect to db failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK")

	poolConfig.MaxConns = 5

	conn, err := domens.NewConnection(poolConfig)

	_, err = conn.Exec(context.Background(), ";")

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

	select {}
}
