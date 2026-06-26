package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DashboardRouter struct {
	server *gin.Engine
	db     *gorm.DB
}

func NewDashboardRouter(server *gin.Engine, db *gorm.DB) *DashboardRouter {
	return &DashboardRouter{
		server: server,
		db:     db,
	}
}

func (dr *DashboardRouter) setupDashboardRoutes() {
	dashboard := dr.server.Group("/dashboard")
	{
		dashboard.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the Dashboard",
			})
		})
	}
}

func (dr *DashboardRouter) DashboardRoutes() {
	dr.setupDashboardRoutes()
}

func SetupDashboardRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewDashboardRouter(server, db)
	router.DashboardRoutes()
}
