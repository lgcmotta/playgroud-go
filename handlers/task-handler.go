package handlers

import (
	"github.com/lgcmotta/playground/domain"
)

type TaskModel struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ITaskHandler interface {
	GetById(id int) *TaskModel
	Add(taskModel TaskModel)
}

type TaskHandler struct {
	repository domain.ITaskRepository
}

func (handler *TaskHandler) GetById(id int) *TaskModel {
	task := handler.repository.GetById(id)
	return mapTask(task)
}

func mapTask(task *domain.Task) *TaskModel {
	taskModel := new(TaskModel)
	taskModel.Id = task.GetId()
	taskModel.Title = task.GetTitle()
	taskModel.Completed = task.IsCompleted()
	return taskModel
}

func (handler *TaskHandler) Add(taskModel TaskModel) {
	task := domain.NewTask(taskModel.Id, taskModel.Title, taskModel.Completed)
	handler.repository.Add(*task)
}

func NewTaskHandler(repository domain.ITaskRepository) ITaskHandler {
	handler := new(TaskHandler)
	handler.repository = repository
	return handler
}
