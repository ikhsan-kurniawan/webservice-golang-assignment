package controllers

import (
	"mygram/models"
	"mygram/repository"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type photoController struct {
	photoRepository repository.IPhotoRepository
}

func NewPhotoController(photoRepository repository.IPhotoRepository) *photoController {
	return &photoController{
		photoRepository: photoRepository,
	}
}

func (pc *photoController) CreatePhoto(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	photo := models.Photo{}

	err := ctx.ShouldBindJSON(&photo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}

	photo.UserID = uint(userID)

	createdPhoto, err := pc.photoRepository.Create(photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdPhoto)
}

func (pc *photoController) GetPhotos(ctx *gin.Context) {
	photos, err := pc.photoRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func (pc *photoController) GetPhotoById(ctx *gin.Context) {
	photoID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	photo, err := pc.photoRepository.GetOne(photoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func (pc *photoController) UpdatePhoto(ctx *gin.Context) {
	photoID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var photo models.Photo
	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	photo.ID = uint(photoID)

	updatedPhoto, err := pc.photoRepository.Update(photo, photoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedPhoto)
}

func (pc *photoController) DeletePhoto(ctx *gin.Context) {
	photoID, err := strconv.Atoi(ctx.Param("photoId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = pc.photoRepository.Delete(photoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Your photo has been successfully deleted",
	})
}