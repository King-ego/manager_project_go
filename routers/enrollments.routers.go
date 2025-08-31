package routers

import (
	"manager_project/controllers"
	"manager_project/repositories"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EnrollmentsRouters struct {
	server *gin.Engine
	db     *gorm.DB
}

func NewEnrollmentsRoutes(server *gin.Engine, db *gorm.DB) *EnrollmentsRouters {
	return &EnrollmentsRouters{
		server: server,
		db:     db,
	}
}

func (er *EnrollmentsRouters) setupEnrollmentsRoutes() {
	repository := repositories.NewEnrollmentsRepository(er.db)

	useCase := usecases.NewEnrollmentsUseCase(repository)

	controller := controllers.NewEnrollmentsController(useCase)

	enrollments := er.server.Group("/enrollments")

	{
		enrollments.GET("/", controller.CreateEnrollment)
	}

}

func (er *EnrollmentsRouters) EnrollmentsRoutes() {
	er.setupEnrollmentsRoutes()
}

func SetupEnrollmentsRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewEnrollmentsRoutes(server, db)
	router.EnrollmentsRoutes()
}
