package routers

import (
	"manager_project/controllers"
	"manager_project/repositories"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StudentsRouters struct {
	server *gin.Engine
	db     *gorm.DB
}

func NewStudentsRouters(server *gin.Engine, db *gorm.DB) *StudentsRouters {
	return &StudentsRouters{
		server: server,
		db:     db,
	}
}

func (r *StudentsRouters) setupStudentsRouters() {
	studentRepository := repositories.NewUserRepository(r.db)
	studentsUseCase := usecases.NewStudentsUseCase(studentRepository)
	usersController := controllers.NewUserController(studentsUseCase)

	users := r.server.Group("/students")
	{

		users.POST("/", usersController.CreateUser)
	}
}

func (r *StudentsRouters) StudentsRouters() {
	r.setupStudentsRouters()
}

func SetupStudentsRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewStudentsRouters(server, db)
	router.StudentsRouters()
}
