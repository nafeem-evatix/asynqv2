package main

import (
	"fmt"

	"github.com/hibiken/asynq"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379", // Redis server address
	}

	inspector := asynq.NewInspector(redisConnection)
	defer inspector.Close()

	servers, err := inspector.Servers()
	if err != nil {
		panic(err)
	}

	fmt.Println(servers)
}
