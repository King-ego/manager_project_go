package routers

import "github.com/gin-gonic/gin"

type TasksRouter struct {
	server *gin.Engine
}

func NewTasksRouter(server *gin.Engine) *TasksRouter {
	return &TasksRouter{
		server: server,
	}
}

func (r *TasksRouter) setupTaskRoutes() {

}

func (r *TasksRouter) Routers() {
	r.setupTaskRoutes()
}

func SetupTasksRoutes(server *gin.Engine) {
	tasksRouter := NewTasksRouter(server)
	tasksRouter.Routers()
}
