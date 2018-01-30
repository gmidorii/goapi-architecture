package form

import (
	"errors"
	"fmt"
	"net/http"
)

// Task is task handler parameter.
type Task struct {
	Id     string   `json:"id"`
	Labels []string `json:"label"`
}

// valid method check parameter valid.
// must put `return true` bottom
func (f Task) valid() (bool, error) {
	if f.Id == "" {
		return false, errors.New("lack `id` query param")
	}

	return true, nil
}

func NewTask(r *http.Request) (Task, error) {
	vm := r.URL.Query()
	f := Task{
		Id:     vm.Get("id"),
		Labels: dotParse(vm.Get("label")),
	}

	if b, err := f.valid(); !b {
		return Task{}, errors.New(fmt.Sprintf("validation error: %v", err))
	}

	return f, nil
}
