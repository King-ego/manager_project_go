package models

type Score struct {
	StudentID string  `json:"student_id"`
	Subject   string  `json:"subject"`
	Score     float64 `json:"score"`
}
