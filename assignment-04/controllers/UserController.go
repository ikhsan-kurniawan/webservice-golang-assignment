package controllers

import (
	"mygram/models"
	"mygram/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepository repository.IUserRepository
}

func NewUserController(userRepository repository.IUserRepository) *userController {
	return &userController{
		userRepository: userRepository,
	}
}

func (uc *userController) UserRegister(ctx *gin.Context) {
	var newUser models.User

	err := ctx.ShouldBindJSON(&newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
			"message": err.Error(),
		})
		return
	}

	registeredUser, err := uc.userRepository.Register(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": "berhasil",
		"message": registeredUser,
	})
}