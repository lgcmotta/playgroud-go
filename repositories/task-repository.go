package repositories

import (
	"github.com/lgcmotta/playground/domain"
	"golang.org/x/exp/slices"
)

type TaskRepository struct {
	tasks []domain.Task
}

func (repository *TaskRepository) GetById(id int) *domain.Task {
	index := slices.IndexFunc(repository.tasks, func(task domain.Task) bool {
		taskId := task.GetId()
		return taskId == id
	})

	return &repository.tasks[index]
}

func (repository *TaskRepository) Add(task domain.Task) {
	repository.tasks = append(repository.tasks, task)
}

func NewTaskRepository() domain.ITaskRepository {
	repository := new(TaskRepository)
	return repository
}
