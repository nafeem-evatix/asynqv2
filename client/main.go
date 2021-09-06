package main

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/nafeem-evatix/asynqv2/server/adder"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379", // Redis server address
	}

	// Create a new Asynq client.
	client := asynq.NewClient(redisConnection)
	defer client.Close()

	adder := adder.NewAdder(10, 10, &redisConnection)
	task, _ := adder.GetTask()

	_, err := client.Enqueue(
		task, // task payload
		asynq.Queue("queue2"),
		asynq.MaxRetry(5),
	)

	if err != nil {
		log.Fatal(err)
	}

}
