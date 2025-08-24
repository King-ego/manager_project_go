package router

import (
	"manager_project/routers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SetupRouter struct {
	server *gin.Engine
	db     *gorm.DB
}

func NewRouter(server *gin.Engine, db *gorm.DB) *SetupRouter {
	return &SetupRouter{
		server: server,
		db:     db,
	}
}

func (r *SetupRouter) setupRouters() {
	routers.SetupStudentsRoutes(r.server, r.db)
	routers.SetupTeacherRoutes(r.server, r.db)
	routers.SetupScoreRoutes(r.server, r.db)
	routers.SetupEnrollmentsRoutes()
}

func (r *SetupRouter) Routers() {
	r.setupRouters()
}

func SetupAllRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewRouter(server, db)
	router.Routers()
}
