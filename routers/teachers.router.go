package routers

import (
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
	teachers := r.server.Group("/teachers")
	{
		teachers.POST("/")
		teachers.GET("/:teacherId")
		teachers.GET("/")
	}

}

func (r *TeacherRouter) TeacherRouters() {
	r.setupTeacherRouters()
}

func SetupTeacherRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewTeacherRouter(server, db)
	router.TeacherRouters()
}
