package form

import (
	"errors"
	"fmt"
	"net/http"
)

// TaskForm is task handler parameter.
type TaskForm struct {
	Id     string   `json:"id"`
	Labels []string `json:"label"`
}

// valid method check parameter valid.
// must put `return true` bottom
func (f TaskForm) valid() (bool, error) {
	if f.Id == "" {
		return false, errors.New("lack `id` query param")
	}

	return true, nil
}

func NewTask(r *http.Request) (TaskForm, error) {
	vm := r.URL.Query()
	f := TaskForm{
		Id:     vm.Get("id"),
		Labels: dotParse(vm.Get("label")),
	}

	if b, err := f.valid(); !b {
		return TaskForm{}, errors.New(fmt.Sprintf("validation error: %v", err))
	}

	return f, nil
}
