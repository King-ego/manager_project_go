package controllers

import (
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	useCase usecases.ClassesUseCase
}

func NewClassesController(useCase usecases.ClassesUseCase) *ClassesController {
	return &ClassesController{
		useCase: useCase,
	}
}

func (cc *ClassesController) CreateClasses(c *gin.Context) {
	var createClasses usecases.ClassesUseCase

	if err := createClasses.CreateClasses(); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"data": "Classes created successfully",
	})
}
