package dto

type CreateScoreDTO struct {
	Comment        string  `json:"comment"`
	ScoreValue     float64 `json:"score_value" binding:"required"`
	Weight         float64 `json:"weight" binding:"required"`
	AssessmentDate string  `json:"assessment_date"`
	AssessmentType string  `json:"assessment_type" binding:"required"`
	EnrollmentID   string  `json:"enrollment_id" binding:"required"`
}
