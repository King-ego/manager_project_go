package usecases

import (
	"errors"
	"manager_project/models"
	"manager_project/repositories"
	"time"
)

type StudentsUseCase interface {
	Create(name string, email string, password string, birthDate string) (error, string)
}

type studentsUseCase struct {
	repository repositories.StudentsRepository
}

func NewStudentsUseCase(repository repositories.StudentsRepository) StudentsUseCase {
	return &studentsUseCase{
		repository: repository,
	}
}

func (uc *studentsUseCase) Create(name string, email string, password string, birthDate string) (error, string) {
	if len(name) < 3 {
		return errors.New("name must be at least 3 characters"), ""
	}

	parsedDate, err := time.Parse("2006-01-02", birthDate)

	if err != nil {
		return errors.New("date must be in YYYY-MM-DD format"), ""
	}

	student := &models.Students{
		Name:      name,
		Email:     email,
		Password:  password,
		BirthDate: parsedDate,
	}

	if err := uc.repository.Save(student); err != nil {
		return err, "failed to save student"
	}

	return nil, "student created successfully"

}
