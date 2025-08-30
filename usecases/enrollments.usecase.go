package usecases

import (
	"manager_project/dto"
	"manager_project/repositories"
)

type EnrollmentsUseCase struct {
	repository repositories.EnrollmentsRepository
}

func NewEnrollmentsUseCase(repository repositories.EnrollmentsRepository) *EnrollmentsUseCase {
	return &EnrollmentsUseCase{
		repository: repository,
	}
}

func (e *EnrollmentsUseCase) CreateEnrollment(enrollment dto.CreateEnrollmentsDTO) error {
	err := e.repository.CreateEnrollment(enrollment)
	if err != nil {
		return nil
	}
	return err
}
