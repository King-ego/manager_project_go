package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	var users = [3]string{"User1", "User2", "User3"}
	ctx.JSON(200, gin.H{
		"message": "Users retrieved successfully",
		"data":    users,
	})

}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var newUser CreateUserRequest
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data":    newUser,
	})
}
