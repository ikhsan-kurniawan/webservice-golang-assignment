package controllers

import (
	"mygram/models"
	"mygram/repository"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type socialMediaController struct {
	socialMediaRepository repository.ISocialMediaRepository
}

func NewSocialMediaController(socialMediaRepository repository.ISocialMediaRepository) *socialMediaController {
	return &socialMediaController{
		socialMediaRepository: socialMediaRepository,
	}
}

func (smc *socialMediaController) CreateSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	socialMedia := models.SocialMedia{}

	err := ctx.ShouldBindJSON(&socialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}

	socialMedia.UserID = uint(userID)

	createdSocialMedia, err := smc.socialMediaRepository.Create(socialMedia)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdSocialMedia)
}

func (smc *socialMediaController) GetSocialMedias(ctx *gin.Context) {
	socialMedias, err := smc.socialMediaRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
}

func (smc *socialMediaController) GetSocialMediaById(ctx *gin.Context) {
	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	socialMedia, err := smc.socialMediaRepository.GetOne(socialMediaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func (smc *socialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var socialMedia models.SocialMedia
	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	socialMedia.ID = uint(socialMediaID)

	updatedSocialMedia, err := smc.socialMediaRepository.Update(socialMedia, socialMediaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedSocialMedia)
}

func (smc *socialMediaController) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = smc.socialMediaRepository.Delete(socialMediaID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Your Social Media has been successfully deleted",
	})
}