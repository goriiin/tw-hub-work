package db

// sudo docker-compose up -d
// sudo docker-compose exec pgdb psql -U postgres -c 'CREATE DATABASE twit_hub'

func Test() {

	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "connect to db failed: %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println("Connection OK")
	//
	//_, err = conn.Exec(context.Background(), ";")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
	//	os.Exit(1)
	//}
	//fmt.Println("Ping OK!")
	//for i := 0; i < 5; i++ {
	//	go func(count int) {
	//		_, err = conn.Exec(context.Background(), ";")
	//		if err != nil {
	//			fmt.Fprintf(os.Stderr, "ping failed: %v\n", err)
	//			os.Exit(1)
	//		}
	//		fmt.Println(count, "Query OK!")
	//		fmt.Printf("connections - MAX: %d, "+
	//			"Iddle: %d, "+
	//			"Total: %d \n",
	//			conn.Stat().MaxConns(),
	//			conn.Stat().IdleConns(),
	//			conn.Stat().TotalConns())
	//	}(i)
	//}
	////conn.Close()
	//select {}
}
