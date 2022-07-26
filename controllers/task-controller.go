package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lgcmotta/playground/handlers"
)

type TaskController struct {
	handler handlers.ITaskHandler
}

func (controller *TaskController) GetById(context *gin.Context) {
	id, _ := strconv.Atoi(context.Param("id"))
	taskModel := controller.handler.GetById(id)
	context.IndentedJSON(http.StatusOK, taskModel)
}

func (controller *TaskController) Add(context *gin.Context) {
	taskModel := new(handlers.TaskModel)
	context.BindJSON(taskModel)
	controller.handler.Add(*taskModel)
}

func NewTaskController(handler handlers.ITaskHandler) *TaskController {
	return &TaskController{handler: handler}
}
