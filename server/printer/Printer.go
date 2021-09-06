package printer

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TaskName = "printer"

type Printer struct {
	Message string
}

func (it *Printer) GetPrinterTask() (*asynq.Task, error) {
	payload, err := json.Marshal(it)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskName, payload), nil
}

func NewPrinter(message string) *Printer {
	return &Printer{
		Message: message,
	}
}
