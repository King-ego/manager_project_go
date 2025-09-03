package repositories

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type ClassesRepository interface {
	CreateClasses(classes *models.Classes) error
	GetClassesByStudentId(classesId string) (*models.Classes, error)
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

func (er *classesRepository) GetClassesByStudentId(classesId string) (*models.Classes, error) {
	var classes models.Classes
	if err := er.db.Where("id = ?", classesId).First(&classes).Error; err != nil {
		return nil, err
	}
	return &classes, nil
}
