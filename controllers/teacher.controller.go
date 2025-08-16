package controllers

import (
	"manager_project/dto"

	"github.com/gin-gonic/gin"
)

type TeacherController struct{}

func NewTeacherController() *TeacherController {
	return &TeacherController{}
}

func (t *TeacherController) CreateTeacher(c *gin.Context) {
	var createTeacherDTO dto.CreateTeacherDTO

	if err := c.ShouldBindJSON(&createTeacherDTO); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	text := "Teacher created successfully"

	c.JSON(201, gin.H{"data": text})
}
