package dto

import "manager_project/models"

func (dto *CreateEnrollmentsDTO) ToEnrollmentModel() *models.Enrollment {
	return &models.Enrollment{
		StudentID: dto.StudentID,
		ClassID:   dto.ClassID,
		Status:    dto.Status,
	}
}

type CreateEnrollmentsDTO struct {
	StudentID string `json:"student_id" binding:"required"`
	ClassID   string `json:"class_id" binding:"required"`
	Status    bool   `json:"status"`
}
