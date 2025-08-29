package repositories

import (
	"fmt"
	"manager_project/models"

	"gorm.io/gorm"
)

type EnrollmentsRepository interface {
	CreateEnrollment(enrollment models.Enrollment) error
}

type enrollmentsRepository struct {
	db *gorm.DB
}

func NewEnrollmentsRepository(db *gorm.DB) EnrollmentsRepository {
	return &enrollmentsRepository{db: db}
}

func (r *enrollmentsRepository) CreateEnrollment(enrollment models.Enrollment) error {
	fmt.Println("enrollment in repo:", enrollment)
	if err := r.db.Create(&enrollment).Error; err != nil {
		return err
	}
	return nil
}
