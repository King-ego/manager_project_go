package controllers

import "github.com/gin-gonic/gin"

type TaskController struct {
}

func NewTaskController() TaskController {
	return TaskController{}
}

func (tc *TaskController) GetTasks(ctx *gin.Context) {
	var tasks = [3]string{"Task1", "Task2", "Task3"}
	ctx.JSON(200, gin.H{
		"message": "Tasks retrieved successfully",
		"data":    tasks,
	})
}
