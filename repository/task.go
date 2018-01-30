package repository

import "time"

type Task struct {
	Id    string
	Name  string
	Out   time.Time
	Label string
}

func FetchTaskById(id string) (Task, error) {
	// sql package
	// ...

	return Task{
		Id:    "task1",
		Name:  "create arch",
		Out:   time.Now(),
		Label: "blue",
	}, nil
}
