package controllers

import (
	"manager_project/dto"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
)

type EnrollmentsController struct {
	enrollmentsUseCase *usecases.EnrollmentsUseCase
}

func NewEnrollmentsController(useCase *usecases.EnrollmentsUseCase) *EnrollmentsController {
	return &EnrollmentsController{
		enrollmentsUseCase: useCase,
	}
}

func (ec *EnrollmentsController) CreateEnrollment(c *gin.Context) {
	enrollment := dto.CreateEnrollmentsDTO{}

	if err := c.ShouldBindJSON(&enrollment); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := ec.enrollmentsUseCase.CreateEnrollment(enrollment)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(201, gin.H{
		"data": "Enrollment created successfully",
	})

}
