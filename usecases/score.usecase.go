package usecases

import (
	"manager_project/models"
	"manager_project/repositories"
)

type ScoreUseCase struct {
	repository repositories.ScoreRepository
}

func NewScoreUseCase(repository repositories.ScoreRepository) *ScoreUseCase {
	return &ScoreUseCase{
		repository: repository,
	}
}

func (s *ScoreUseCase) CreateScore() (string, error) {
	err := s.repository.CreateScore(&models.Score{})
	if err != nil {
		return "", err
	}
	return "", nil
}

func (s *ScoreUseCase) GetByStudentID(studentId string) ([]models.Score, error) {
	scores, err := s.repository.GetByStudentID(studentId)
	if err != nil {
		return nil, err
	}
	return scores, nil
}
