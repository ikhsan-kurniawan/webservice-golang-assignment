package controllers

import (
	"mygram/helpers"
	"mygram/models"
	"mygram/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserDTO struct {
	ID			uint    	`json:"id"`
	Username    string		`json:"username"`
	Email       string		`json:"email"`
	Age         int			`json:"age"`
	CreatedAt 	*time.Time	`json:"created_at"`
	UpdatedAt 	*time.Time	`json:"updated_at"`
}

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
			"error": "Bad Request",
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

	response := UserDTO {
		ID: registeredUser.ID,
		Username: registeredUser.Username,
		Email: registeredUser.Email,
		Age: registeredUser.Age,
		CreatedAt: registeredUser.CreatedAt,
		UpdatedAt: registeredUser.UpdatedAt,
	}

	ctx.JSON(http.StatusCreated, response)
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

	response := UserDTO {
		ID: updatedUser.ID,
		Username: updatedUser.Username,
		Email: updatedUser.Email,
		Age: updatedUser.Age,
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, response)
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
