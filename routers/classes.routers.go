package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ClassesRouter struct {
	server *gin.Engine
	db     *gorm.DB
}

func NewClassesRouter(server *gin.Engine, db *gorm.DB) *ClassesRouter {
	return &ClassesRouter{
		server: server,
		db:     db,
	}
}

func (pr *ClassesRouter) setupClassesRouter() {
	classes := pr.server.Group("/classes")
	{
		classes.GET("/")
	}
}

func (pr *ClassesRouter) ClassesRouter() {
	pr.setupClassesRouter()
}

func SetupClassesRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewClassesRouter(server, db)
	router.ClassesRouter()
}
