package model

type StudentGrades struct {
	StudentID string `json:"student_id"`
	Year      int    `json:"year"`
	Semester  int    `json:"semester"`
}
