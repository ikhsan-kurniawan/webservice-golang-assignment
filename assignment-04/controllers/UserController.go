package controllers

import (
	"mygram/helpers"
	"mygram/models"
	"mygram/repository"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
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
		"id": registeredUser.ID,
		"username": registeredUser.Username,
		"email": registeredUser.Email,
		"password": registeredUser.Password,
	})
}

func (uc *userController) UserLogin(ctx *gin.Context) {
	user := models.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "bad request",
			"message": err.Error(),
		})
		return
	}

	loginUser, err := uc.userRepository.Login(user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
			"message": err.Error(),
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(loginUser.Password), []byte(user.Password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token, err := helpers.GenerateToken(loginUser.ID, loginUser.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Parsed Token Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": loginUser.ID,
		"email": loginUser.Email,
		"token": token,
	})
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	user := models.User{}
	userID, err := strconv.Atoi(ctx.Param("userId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}

	updatedUser, err := uc.userRepository.Update(user, userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id": updatedUser.ID,
		"email": updatedUser.Email,
		"username": updatedUser.Username,
		"age": updatedUser.Age,
		"updated_at": updatedUser.UpdatedAt,
	})
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))

	err := uc.userRepository.Delete(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Your account has been successfully deleted",
	})
}
