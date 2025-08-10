package routers

import "github.com/gin-gonic/gin"

type UsersRouters struct {
	server *gin.Engine
}

func NewUsersRouters(server *gin.Engine) *UsersRouters {
	return &UsersRouters{
		server: server,
	}
}

func (r *UsersRouters) setupUsersRouters() {
	users := r.server.Group("/users")
	{
		users.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "List of users",
			})
		})
		users.POST("/", func(c *gin.Context) {
			c.JSON(201, gin.H{
				"message": "User created",
			})
		})
	}
}

func (r *UsersRouters) UsersRouters() {
	r.setupUsersRouters()
}

func SetupUsersRoutes(server *gin.Engine) {
	router := NewUsersRouters(server)
	router.UsersRouters()
}
