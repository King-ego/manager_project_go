package routers

import (
	"manager_project/controllers"
	"manager_project/repositories"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ScoreRouters struct {
	db     *gorm.DB
	server *gin.Engine
}

func NewScoreRouters(server *gin.Engine, db *gorm.DB) *ScoreRouters {
	return &ScoreRouters{
		db:     db,
		server: server,
	}
}

func (r *ScoreRouters) setupScoreRouters() {
	scoreRepository := repositories.NewScoreRepository(r.db)
	scoreUseCase := usecases.NewScoreUseCase(scoreRepository)
	scoreController := controllers.NewScoreController(scoreUseCase)

	scores := r.server.Group("/scores")
	{
		scores.POST("/", scoreController.CreateScore)
		scores.GET("/:studentId")
	}
}

func (r *ScoreRouters) ScoreRouters() {
	r.setupScoreRouters()
}

func SetupScoreRoutes(server *gin.Engine, db *gorm.DB) {
	router := NewScoreRouters(server, db)
	router.ScoreRouters()
}
