package usecases

import (
	"manager_project/models"
	"manager_project/repositories"
)

type ClassesUseCase struct {
	repository repositories.ClassesRepository
}

func NewClassesUseCase(repository repositories.ClassesRepository) *ClassesUseCase {
	return &ClassesUseCase{
		repository: repository,
	}
}

func (c *ClassesUseCase) CreateClasses(createClasses *models.Classes) error {
	err := c.repository.CreateClasses(createClasses)
	if err != nil {
		return err
	}
	return nil
}
