package controllers

import (
	"manager_project/dto"
	"net/http"

	"manager_project/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUseCase usecases.StudentsUseCase
}

func NewUserController(uc usecases.StudentsUseCase) *UserController {
	return &UserController{userUseCase: uc}
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	var createStudentDTO dto.CreateStudentsDTO

	if err := c.ShouldBindJSON(&createStudentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, text := ctrl.userUseCase.Create(createStudentDTO.Name, createStudentDTO.Email, createStudentDTO.Password, createStudentDTO.BirthDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, text)
}
