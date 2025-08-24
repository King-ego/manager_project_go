package controllers

type EnrollmentsController struct{}

func NewEnrollmentsController() *EnrollmentsController {
	return &EnrollmentsController{}
}

func (ec *EnrollmentsController) CreateEnrollment() string {
	return "Enrollment created"
}
