package todo

import (
	"reflect"
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

func (t *Todo) isExisting(task Task) bool {
	for _, todoitem := range *t {
		if todoitem.Id == task.Id && todoitem.TaskTitle == task.TaskTitle &&
			todoitem.Status == task.Status {
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
