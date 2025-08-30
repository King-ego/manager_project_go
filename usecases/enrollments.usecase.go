package usecases

import (
	"manager_project/dto"
	"manager_project/repositories"

	"golang.org/x/crypto/openpgp/errors"
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
	enrollmentModel := enrollment.ToEnrollmentModel()
	if enrollmentModel == nil {
		return errors.InvalidArgumentError("Invalid enrollment data")
	}

	err := e.repository.CreateEnrollment(enrollmentModel)
	if err != nil {
		return err
	}
	return err
}
