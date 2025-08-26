package repositories

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type ClassesRepository interface {
	CreateClasses(classes *models.Classes) error
}
type classesRepository struct {
	db *gorm.DB
}

func NewClassesRepository(db *gorm.DB) ClassesRepository {
	return &classesRepository{db: db}
}

func (er *classesRepository) CreateClasses(classes *models.Classes) error {
	if err := er.db.Create(classes).Error; err != nil {
		return err
	}
	return nil
}
