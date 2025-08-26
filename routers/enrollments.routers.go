package routers

import (
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
	enrollments := er.server.Group("/enrollments")
	{
		enrollments.GET("/")
	}

}

func (er *EnrollmentsRouters) EnrollmentsRoutes() {
	er.setupEnrollmentsRoutes()
}

func SetupEnrollmentsRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewEnrollmentsRoutes(server, db)
	router.EnrollmentsRoutes()
}
