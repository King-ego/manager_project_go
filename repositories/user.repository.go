package repositories

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(student *models.Students) error
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Save(student *models.Students) error {
	return r.db.Create(student).Error
}
