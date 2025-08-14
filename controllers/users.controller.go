package controllers

import (
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
	var request struct {
		Name      string `json:"name" binding:"required"`
		Email     string `json:"email" binding:"required,email"`
		Password  string `json:"password" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required,date=2006-01-02"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err, text := ctrl.userUseCase.Create(request.Name, request.Email, request.Password, request.BirthDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, text)
}
