package routers

import (
	"fmt"
	"manager_project/controllers"

	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	server *gin.Engine
}

func NewUserRouter(server *gin.Engine) *UserRouter {
	return &UserRouter{
		server: server,
	}
}

func (r *UserRouter) setupUserRoutes() {
	fmt.Println("Setting up user routes")
	usersController := controllers.NewUserController()

	r.server.GET("/users", usersController.GetUsers)
	r.server.POST("/users", usersController.CreateUser)
}

func (r *UserRouter) Routers() {
	r.setupUserRoutes()
}

func SetupUserRoutes(server *gin.Engine) {
	userRouter := NewUserRouter(server)
	userRouter.Routers()
}
