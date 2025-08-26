package usecases

import "manager_project/repositories"

type ClassesUseCase struct {
	repository repositories.ClassesRepository
}

func NewClassesUseCase(repository repositories.ClassesRepository) *ClassesUseCase {
	return &ClassesUseCase{
		repository: repository,
	}
}

func (c *ClassesUseCase) CreateClasses() error {
	err := c.repository.CreateEnrollment(nil)
	if err != nil {
		return err
	}
	return nil
}
