package controllers

import "manager_project/usecases"

type ScoreController struct {
	scoreUseCase usecases.ScoreUseCase
}

func NewScoreController(useCase usecases.ScoreUseCase) *ScoreController {
	return &ScoreController{
		scoreUseCase: useCase,
	}
}

func (c *ScoreController) CreateScore() (string, error) {
	text, err := c.scoreUseCase.CreateScore()
	if err != nil {
		return "", err
	}
	return text, nil
}
