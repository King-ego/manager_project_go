package controllers

import (
	"manager_project/dto"
	"net/http"

	"manager_project/usecases"

	"github.com/gin-gonic/gin"
)

type StudentsController struct {
	studentsUseCase *usecases.StudentsUseCase
}

func NewStudentsController(uc *usecases.StudentsUseCase) *StudentsController {
	return &StudentsController{studentsUseCase: uc}
}

func (ctrl *StudentsController) CreateStudents(c *gin.Context) {
	var createStudentDTO dto.CreateStudentsDTO

	if err := c.ShouldBindJSON(&createStudentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, text := ctrl.studentsUseCase.Create(createStudentDTO.Name, createStudentDTO.Email, createStudentDTO.Password, createStudentDTO.BirthDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, text)
}

func (ctrl *StudentsController) GetStudentByID(c *gin.Context) {
	studentId := c.Param("studentId")

	if studentId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student ID is required"})
		return
	}

	student, err := ctrl.studentsUseCase.GetByID(studentId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}
