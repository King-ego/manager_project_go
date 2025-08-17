package repositories

import (
	"manager_project/models"

	"gorm.io/gorm"
)

type TeacherRepository interface {
	Save(teacher *models.Teacher) error
	GetByID(teacherId string) (models.Teacher, error)
}

type teacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) TeacherRepository {
	return &teacherRepository{db: db}
}

func (r *teacherRepository) Save(teacher *models.Teacher) error {
	return r.db.Create(teacher).Error
}

func (r *teacherRepository) GetByID(teacherId string) (models.Teacher, error) {
	var teacher models.Teacher
	err := r.db.Where("id = ?", teacherId).First(&teacher).Error
	if err != nil {
		return models.Teacher{}, err
	}
	return teacher, nil
}
