package routers

import (
	"manager_project/controllers"
	"manager_project/repositories"
	"manager_project/usecases"

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
	classesRepository := repositories.NewClassesRepository(pr.db)
	classesUseCase := usecases.NewClassesUseCase(classesRepository)
	classesController := controllers.NewClassesController(classesUseCase)

	classes := pr.server.Group("/classes")
	{
		classes.GET("/", classesController.CreateClasses)
	}
}

func (pr *ClassesRouter) ClassesRouter() {
	pr.setupClassesRouter()
}

func SetupClassesRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewClassesRouter(server, db)
	router.ClassesRouter()
}
