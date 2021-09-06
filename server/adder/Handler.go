package adder

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"github.com/nafeem-evatix/asynqv2/server/printer"
)

func Handler(ctx context.Context, task *asynq.Task) error {
	var it Adder
	if err := json.Unmarshal(task.Payload(), &it); err != nil {
		return err
	}

	result := it.Num1 + it.Num2
	printer := printer.NewPrinter(fmt.Sprintf("%v+%v=%v", it.Num1, it.Num2, result))

	printerTask, err := printer.GetPrinterTask()
	if err != nil {
		return err
	}

	client := asynq.NewClient(it.RedisClientOpt)
	_, errEnqueue := client.Enqueue(printerTask, asynq.Queue("queue1"))
	if errEnqueue != nil {
		return errEnqueue
	}

	defer client.Close()

	return nil
}
