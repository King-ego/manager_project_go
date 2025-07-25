package routers

import (
	"manager_project/controllers"

	"github.com/gin-gonic/gin"
)

type TasksRouter struct {
	server *gin.Engine
}

func NewTasksRouter(server *gin.Engine) *TasksRouter {
	return &TasksRouter{
		server: server,
	}
}

func (r *TasksRouter) setupTaskRoutes() {
	taskController := controllers.NewTaskController()
	r.server.GET("/tasks", taskController.GetTasks)

}

func (r *TasksRouter) Routers() {
	r.setupTaskRoutes()
}

func SetupTasksRoutes(server *gin.Engine) {
	tasksRouter := NewTasksRouter(server)
	tasksRouter.Routers()
}
