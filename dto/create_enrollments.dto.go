package dto

type CreateEnrollmentsDTO struct {
	StudentID string `json:"student_id" binding:"required"`
	ClassID   string `json:"class_id" binding:"required"`
	Status    *bool  `json:"status"`
}
