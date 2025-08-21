package repositories

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type ScoreRepository interface {
	CreateScore(score *models.Score) error
}

type scoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) ScoreRepository {
	return &scoreRepository{db: db}
}

func (r *scoreRepository) CreateScore(score *models.Score) error {
	if err := r.db.Create(score).Error; err != nil {
		return err
	}
	return nil
}
