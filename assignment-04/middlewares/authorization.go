package middlewares

import (
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID, err := strconv.Atoi(ctx.Param("userId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad request",
				"message": err.Error(),
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)

		if userID != int(userData["id"].(float64)) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "kamu ga boleh akses data ini",
			})
			return
		}
		ctx.Next()
	}
}

func PhotoAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		photoID, err := strconv.Atoi(ctx.Param("photoId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
		
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		var photo models.Photo
		if err := db.First(&photo, photoID).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error":   "Internal Server Error",
				"message": err.Error(),
			})
			return
		}

		if userID != int(photo.UserID) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "kamu ga boleh akses data ini",
			})
			return
		}
		ctx.Next()
	}
}

func CommentAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		commentID, err := strconv.Atoi(ctx.Param("commentId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
		
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		var comment models.Comment
		if err := db.First(&comment, commentID).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			})
			return
		}

		if userID != int(comment.UserID) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "kamu ga boleh akses data ini",
			})
			return
		}
		ctx.Next()
	}
}

func SocialMediaAuthorization(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		socialMediaID, err := strconv.Atoi(ctx.Param("socialMediaId"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": err.Error(),
			})
			return
		}
		
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := int(userData["id"].(float64))

		var socialMedia models.SocialMedia
		if err := db.First(&socialMedia, socialMediaID).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Bad Request",
				"message": err.Error(),
			})
			return
		}

		if userID != int(socialMedia.UserID) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Unauthorized",
				"message": "kamu ga boleh akses data ini",
			})
			return
		}
		ctx.Next()
	}
}