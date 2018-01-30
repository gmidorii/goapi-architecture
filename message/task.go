package message

import (
	"github.com/midorigreen/goapi-architecture/domain"
)

type Task struct {
	Name   string `json:"name"`
	Charge string `json:"charge"`
}

func ConvertTask(d domain.Task) (Task, error) {
	return Task{
		Name:   d.Name,
		Charge: d.Assign,
	}, nil
}
