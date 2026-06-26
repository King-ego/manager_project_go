package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DashboardRouter struct {
	server *gin.Engine
	db     *gorm.DB
}
