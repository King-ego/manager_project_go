package routers

import (
	"manager_project/controllers"
	"manager_project/repositories"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TeacherRouter struct {
	server *gin.Engine
	db     *gorm.DB
}

func NewTeacherRouter(server *gin.Engine, db *gorm.DB) *TeacherRouter {
	return &TeacherRouter{
		server: server,
		db:     db,
	}

}

func (r *TeacherRouter) setupTeacherRouters() {
	teacherRepository := repositories.NewTeacherRepository(r.db)
	teacherUseCase := usecases.NewTeacherUseCase(teacherRepository)
	teachersController := controllers.NewTeacherController(teacherUseCase)
	teachers := r.server.Group("/teachers")
	{
		teachers.POST("/", teachersController.CreateTeacher)
		teachers.GET("/:teacherId", teachersController.GetTeacherByID)
	}

}

func (r *TeacherRouter) TeacherRouters() {
	r.setupTeacherRouters()
}

func SetupTeacherRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewTeacherRouter(server, db)
	router.TeacherRouters()
}
