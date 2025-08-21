package repositories

import "gorm.io/gorm"

type ScoreRepository interface {
	CreateScore() error
}

type scoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) ScoreRepository {
	return &scoreRepository{db: db}
}

func (r *scoreRepository) CreateScore() error {
	return nil
}
