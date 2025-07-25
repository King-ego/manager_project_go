package controllers

import "github.com/gin-gonic/gin"

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
