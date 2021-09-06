package middlewares

import (
	"context"
	"fmt"

	"github.com/hibiken/asynq"
)

func Log(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		taskName := t.Type()
		taskId, _ := asynq.GetTaskID(ctx)
		queueName, _ := asynq.GetQueueName(ctx)

		fmt.Printf(
			"Task With Name:%v And Id:%v , From Queue :%v Is Now Being Executed",
			taskName,
			taskId,
			queueName,
		)

		fmt.Println()

		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}

		fmt.Println("Processing End")
		return nil
	})
}