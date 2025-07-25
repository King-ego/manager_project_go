package routers

import "github.com/gin-gonic/gin"

type TasksRouter struct {
	server *gin.Engine
}

func NewProjectRouter(server *gin.Engine) *TasksRouter {
	return &TasksRouter{
		server: server,
	}
}

func (r *TasksRouter) setupProjectRoutes() {

}

func (r *TasksRouter) Routers() {
	r.setupProjectRoutes()
}

func SetupProjectRoutes(server *gin.Engine) {
	projectRouter := NewProjectRouter(server)
	projectRouter.Routers()
}
