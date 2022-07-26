package domain

type ITaskRepository interface {
	GetById(id int) *Task
	Add(task Task)
}

type Task struct {
	id        int
	title     string
	completed bool
}

func (task *Task) GetId() int {
	return task.id
}

func (task *Task) GetTitle() string {
	return task.title
}

func (task *Task) IsCompleted() bool {
	return task.completed
}

func NewTask(id int, title string, completed bool) *Task {
	task := new(Task)
	task.id = id
	task.title = title
	task.completed = completed
	return task
}
