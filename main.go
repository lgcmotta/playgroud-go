package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lgcmotta/playground/controllers"
	"github.com/lgcmotta/playground/handlers"
	"github.com/lgcmotta/playground/repositories"
	"go.uber.org/dig"
)

func main() {
	container := dig.New()
	container.Provide(repositories.NewTaskRepository)
	container.Provide(handlers.NewTaskHandler)
	container.Provide(controllers.NewTaskController)

	router := gin.Default()
	router.GET("/api/tasks/:id", func(ctx *gin.Context) {
		err := container.Invoke(func(controller *controllers.TaskController) {
			controller.GetById(ctx)
		})
		if err == nil {
			panic(err)
		}
	})
	router.POST("/api/tasks", func(ctx *gin.Context) {
		err := container.Invoke(func(controller *controllers.TaskController) {
			controller.Add(ctx)
		})
		if err == nil {
			panic(err)
		}
	})
	router.Run()
}
