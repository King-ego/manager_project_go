package dto

type CreateTeacherDTO struct {
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Specialization string `json:"specialization" binding:"required"`
	HireDate       string `json:"hire_date" binding:"required,datetime=2006-01-02"`
}
