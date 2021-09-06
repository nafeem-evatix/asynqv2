package printer

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

func Handler(ctx context.Context, task *asynq.Task) error {
	var p Printer
	if err := json.Unmarshal(task.Payload(), &p); err != nil {
		return err
	}

	fmt.Println(p.Message)
	return nil
}
