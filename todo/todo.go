package todo

import (
	"slices"
)

type TaskStatus int
type TodoErr string

const (
	NotStarted TaskStatus = iota
	InProgress
	Done
)

const (
	AlreadyPresent = TodoErr("could not add this task because it already exists")
	IdConflict     = TodoErr("the task id is already stored in the todo")
	NotFound       = TodoErr("the task is not registered in the todo list")
)

func (e TodoErr) Error() string {
	return string(e)
}

type Task struct {
	Id        int
	TaskTitle string
	Status    TaskStatus
}

type Todo []Task

func (t *Task) isEqual(task *Task) bool {
	return t.Id == task.Id && t.Status == task.Status &&
		t.TaskTitle == task.TaskTitle
}

func (t *Todo) isExisting(task Task) bool {
	for _, todoitem := range *t {
		if todoitem.isEqual(&task) {
			return true
		}
	}
	return false
}

func (t *Todo) isConflictingId(task Task) bool {
	for _, todoitem := range *t {
		if todoitem.Id == task.Id {
			return true
		}
	}
	return false
}

func (t *Todo) Add(task Task) error {
	if t.isExisting(task) {
		return AlreadyPresent
	}
	if t.isConflictingId(task) {
		return IdConflict
	}
	if task.Id == 0 {
		if len(*t) == 0 {
			task.Id = 1
		} else {
			task.Id = (*t)[len(*t)-1].Id + 1
		}
	}
	*t = append(*t, task)
	return nil
}

func (t *Todo) Remove(task Task) error {
	if !t.isExisting(task) {
		return NotFound
	}

	for i, todoitem := range *t {
		if todoitem.isEqual(&task) {
			*t = slices.Delete(*t, i, i+1)
		}
	}
	return nil
}
