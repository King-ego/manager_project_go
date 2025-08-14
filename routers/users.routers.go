package routers

import (
	"manager_project/controllers"
	"manager_project/repositories"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UsersRouters struct {
	server *gin.Engine
	db     *gorm.DB
}

func NewUsersRouters(server *gin.Engine, db *gorm.DB) *UsersRouters {
	return &UsersRouters{
		server: server,
		db:     db,
	}
}

func (r *UsersRouters) setupUsersRouters() {
	studentRepository := repositories.NewUserRepository(r.db)
	studentsUseCase := usecases.NewStudentsUseCase(studentRepository)
	usersController := controllers.NewUserController(studentsUseCase)

	users := r.server.Group("/users")
	{

		users.POST("/", usersController.CreateUser)
	}
}

func (r *UsersRouters) UsersRouters() {
	r.setupUsersRouters()
}

func SetupUsersRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewUsersRouters(server, db)
	router.UsersRouters()
}
