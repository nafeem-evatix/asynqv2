package adder

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

const TaskName = "adder"

type Adder struct {
	Num1, Num2     int
	RedisClientOpt *asynq.RedisClientOpt
}

func (it *Adder) GetTask() (*asynq.Task, error) {
	payload, err := json.Marshal(it)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskName, payload), nil
}

func NewAdder(
	num1,
	num2 int,
	redisClientOpt *asynq.RedisClientOpt,
) *Adder {
	return &Adder{
		Num1:           num1,
		Num2:           num2,
		RedisClientOpt: redisClientOpt,
	}
}
