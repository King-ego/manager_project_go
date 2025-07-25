package routers

import "github.com/gin-gonic/gin"

type ProjectRouter struct {
	server *gin.Engine
}

func NewProjectRouter(server *gin.Engine) *ProjectRouter {
	return &ProjectRouter{
		server: server,
	}
}

func (r *ProjectRouter) setupProjectRoutes() {

}

func (r *ProjectRouter) Routers() {
	r.setupProjectRoutes()
}

func SetupProjectRoutes(server *gin.Engine) {
	projectRouter := NewProjectRouter(server)
	projectRouter.Routers()
}
