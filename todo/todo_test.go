package todo

import (
	"reflect"
	"testing"
)

func TestAddTask(t *testing.T) {

	task1 := Task{0, "test", InProgress}
	tmp := &Todo{}
	tmp.Add(task1)
	got := *tmp
	want := Todo{{0, "test", InProgress}}

	if !reflect.DeepEqual(got, want) {
		t.Logf("got: %#v, want: %#v", got, want)

		t.Errorf("got %v, want %v", got, want)
	}
}
