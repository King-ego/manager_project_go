package usecases

import (
	"errors"
	"manager_project/models"
	"manager_project/repositories"
	"time"
)

type TeacherUseCase interface {
	Save(name string, email string, specialization string, hireDate string) error
}

type teacherUseCase struct {
	teacherRepository repositories.TeacherRepository
}

func NewTeacherUseCase(teacherRepository repositories.TeacherRepository) TeacherUseCase {
	return &teacherUseCase{teacherRepository: teacherRepository}
}

func (u *teacherUseCase) Save(name string, email string, specialization string, hireDate string) error {
	parsedDate, err := time.Parse("2006-01-02", hireDate)
	if err != nil {
		return errors.New("date must be in YYYY-MM-DD format")
	}

	teacher := &models.Teacher{
		Name:           name,
		Email:          email,
		Specialization: specialization,
		HireDate:       parsedDate,
	}
	return u.teacherRepository.Save(teacher)
}
