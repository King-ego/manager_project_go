package controllers

import (
	"manager_project/dto"
	"manager_project/usecases"

	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	teacherUseCase usecases.TeacherUseCase
}

func NewTeacherController(uc usecases.TeacherUseCase) *TeacherController {
	return &TeacherController{
		teacherUseCase: uc,
	}
}

func (t *TeacherController) CreateTeacher(c *gin.Context) {
	var createTeacherDTO dto.CreateTeacherDTO

	if err := c.ShouldBindJSON(&createTeacherDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := t.teacherUseCase.Save(createTeacherDTO.Name, createTeacherDTO.Email, createTeacherDTO.HireDate, createTeacherDTO.Specialization)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	text := "Teacher created successfully"

	c.JSON(201, gin.H{"data": text})
}

func (t *TeacherController) GetTeacherByID(c *gin.Context) {
	teacherId := c.Param("teacherId")

	if teacherId == "" {
		c.JSON(400, gin.H{"error": "Teacher ID is required"})
		return
	}

	teacher, err := t.teacherUseCase.GetByID(teacherId)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": teacher})
}
