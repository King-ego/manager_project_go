package router

import (
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
}

func (r *SetupRouter) Routers() {
	r.setupRouters()
}

func SetupAllRoutes(server *gin.Engine) {
	router := NewRouter(server)
	router.Routers()
}
