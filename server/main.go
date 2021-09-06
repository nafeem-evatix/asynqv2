package main

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/nafeem-evatix/asynqv2/server/adder"
	"github.com/nafeem-evatix/asynqv2/server/middlewares"
	"github.com/nafeem-evatix/asynqv2/server/printer"
)

func main() {
	redisConnection := asynq.RedisClientOpt{
		Addr: "localhost:6379", // Redis server address
	}

	server := asynq.NewServer(redisConnection, asynq.Config{
		// Specify how many concurrent workers to use.
		Concurrency: 10,
		Queues: map[string]int{
			"queue1": 4, // processed 40% of the time
			"queue2": 4, // processed 40% of the time
			"queue3": 2, // processed 20% of the time
		},
	})

	mux := asynq.NewServeMux()
	mux.Use(middlewares.Log)

	// Define a task handler for the welcome email task.
	mux.HandleFunc(
		printer.TaskName, // task type
		printer.Handler,  // handler function
	)

	mux.HandleFunc(
		adder.TaskName, // task type
		adder.Handler,  // handler function
	)

	if err := server.Run(mux); err != nil {
		log.Fatal(err)
	}
}
