package repositories

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type StudentsRepository interface {
	Save(student *models.Students) error
	GetByID(studentId string) (models.Students, error)
}
type studentsRepository struct {
	db *gorm.DB
}

func NewStudentsRepository(db *gorm.DB) StudentsRepository {
	return &studentsRepository{db: db}
}

func (r *studentsRepository) Save(student *models.Students) error {
	return r.db.Create(student).Error
}

func (r *studentsRepository) GetByID(studentId string) (models.Students, error) {
	var student models.Students
	err := r.db.Where("id = ?", studentId).First(&student).Error
	if err != nil {
		return models.Students{}, err
	}
	return student, nil
}
