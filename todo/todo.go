package todo

type TaskStatus int

const (
	NotStarted TaskStatus = iota
	InProgress
	Done
)

type Task struct {
	Id     int
	Task   string
	Status TaskStatus
}

type Todo []Task

func (t *Todo) Add(task Task) {
	*t = append(*t, task)
}
