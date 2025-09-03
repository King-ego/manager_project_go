package controllers

import (
	"manager_project/dto"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	useCase *usecases.ClassesUseCase
}

func NewClassesController(useCase *usecases.ClassesUseCase) *ClassesController {
	return &ClassesController{
		useCase: useCase,
	}
}

func (cc *ClassesController) CreateClasses(c *gin.Context) {
	var createClassesDTO dto.CreateClassesDTO
	if err := c.ShouldBindJSON(&createClassesDTO); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	classesModel := createClassesDTO.ToModel()

	if err := cc.useCase.CreateClasses(classesModel); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"data": "Classes created successfully",
	})
}

func (cc *ClassesController) GetClassesByStudentId(c *gin.Context) {
	studentId := c.Param("studentId")
	if studentId == "" {
		c.JSON(400, gin.H{
			"error": "Student ID is required",
		})
		return
	}

	classes, err := cc.useCase.GetClassesByStudentId(studentId)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": classes,
	})
}
