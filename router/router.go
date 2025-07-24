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
	r.server.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}

func (r *SetupRouter) Routers() {
	r.setupRouters()
}

func SetupAllRoutes(server *gin.Engine) {
	router := NewRouter(server)
	router.Routers()
}
