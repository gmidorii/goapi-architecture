package domain

import (
	"time"

	"github.com/midorigreen/goapi-architecture/form"
	"github.com/midorigreen/goapi-architecture/repository"
)

// Task is domain task info.
type Task struct {
	Id       string
	Name     string
	Deadline time.Time
	Assign   string
}

// CreateTaskDomain make task domain struct.
func CreateTaskDomain(f form.Task) (Task, error) {
	repoTask, err := repository.FetchTaskById(f.Id)
	if err != nil {
		return Task{}, err
	}

	repoMember, err := repository.FetchMemberByTaskId(repoTask.Id)
	if err != nil {
		return Task{}, err
	}

	return Task{
		Id:       repoTask.Id,
		Name:     repoTask.Name,
		Deadline: repoTask.Out,
		Assign:   repoMember.Name,
	}, nil
}
