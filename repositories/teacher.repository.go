package repositories

import "gorm.io/gorm"

type TeacherRepository interface {
	Save() error
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
