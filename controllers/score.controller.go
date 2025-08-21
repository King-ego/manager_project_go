package controllers

import (
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
)

type ScoreController struct {
	scoreUseCase *usecases.ScoreUseCase
}

func NewScoreController(useCase *usecases.ScoreUseCase) *ScoreController {
	return &ScoreController{
		scoreUseCase: useCase,
	}
}

func (sc *ScoreController) CreateScore(c *gin.Context) {
	text, err := sc.scoreUseCase.CreateScore()
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(201, gin.H{
		"data": text,
	})
}
