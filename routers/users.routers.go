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
	usersRoutes := controllers.NewUserController()

	r.server.GET("/users", usersRoutes.GetUsers)
}

func (r *UserRouter) Routers() {
	r.setupUserRoutes()
}

func SetupUserRoutes(server *gin.Engine) {
	userRouter := NewUserRouter(server)
	userRouter.Routers()
}
