package hub

import (
	"github.com/midorigreen/goapi-architecture/domain"
	"github.com/midorigreen/goapi-architecture/form"
	"github.com/midorigreen/goapi-architecture/message"
)

func TaskHub(f form.Task) (message.Task, error) {
	// call domain
	// role: fetch data form domain-repository
	d, err := domain.CreateTaskDomain(f)
	if err != nil {
		return message.Task{}, err
	}

	// call message
	// role: convert response type(= message)
	m, err := message.ConvertTask(d)
	if err != nil {
		return message.Task{}, err
	}
	return m, nil
}
