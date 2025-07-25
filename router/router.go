package router

import (
	"manager_project/routers"

	"github.com/gin-gonic/gin"
)

type SetupRouter struct {
	server *gin.Engine
}

func NewRouter(server *gin.Engine) *SetupRouter {
	return &SetupRouter{
		server: server,
	}
}

func (r *SetupRouter) setupRouters() {
	routers.SetupUserRoutes(r.server)
	routers.SetupTasksRoutes(r.server)
}

func (r *SetupRouter) Routers() {
	r.setupRouters()
}

func SetupAllRoutes(server *gin.Engine) {
	router := NewRouter(server)
	router.Routers()
}
