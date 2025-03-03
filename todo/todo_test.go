package todo

import (
	"reflect"
	"testing"
)

func AssertErrors(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestAddTask(t *testing.T) {

	checkAdd := func(t testing.TB, got, want Todo) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	}

	t.Run("Add a new full struct task", func(t *testing.T) {
		task1 := Task{1, "test", InProgress}
		list := &Todo{}
		err := list.Add(task1)
		got := *list
		want := Todo{{1, "test", InProgress}}

		checkAdd(t, got, want)
		AssertErrors(t, err, nil)
	})

	t.Run("Add a task without id", func(t *testing.T) {
		task2 := Task{TaskTitle: "test2", Status: NotStarted}
		list := &Todo{{9, "another task", Done}}
		err := list.Add(task2)
		got := *list
		want := Todo{{9, "another task", Done}, {10, "test2", NotStarted}}

		checkAdd(t, got, want)
		AssertErrors(t, err, nil)
	})

	t.Run("Add a task without completion", func(t *testing.T) {
		taskNotStarted := Task{Id: 12, TaskTitle: "not Started"}
		list := &Todo{}
		err := list.Add(taskNotStarted)
		got := *list
		want := Todo{{12, "not Started", NotStarted}}

		checkAdd(t, got, want)
		AssertErrors(t, err, nil)
	})

	t.Run("Add an existing task", func(t *testing.T) {
		task := Task{Id: 12, TaskTitle: "already exist"}
		list := &Todo{task}
		err := list.Add(task)
		got := *list
		want := Todo{{12, "already exist", NotStarted}}

		checkAdd(t, got, want)
		AssertErrors(t, err, AlreadyPresent)
	})

	t.Run("Add task with an existing id", func(t *testing.T) {
		task := Task{Id: 99, TaskTitle: "id already exist"}
		list := &Todo{{Id: 99, TaskTitle: "another task"}}
		err := list.Add(task)
		got := *list
		want := Todo{{99, "another task", NotStarted}}

		checkAdd(t, got, want)
		AssertErrors(t, err, IdConflict)
	})
}
