package repositories

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Save() error
	GetByID(teacherId string) (models.Teacher, error)
}

type teacherRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

func (r *teacherRepository) Save() error {
	return r.db.Create(&struct{}{}).Error
}

func (r *teacherRepository) GetByID(teacherId string) (models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.Where("id = ?", teacherId).First(&teacher).Error
	if err != nil {
		return models.Teacher{}, err
	}
	return teacher, nil
}
