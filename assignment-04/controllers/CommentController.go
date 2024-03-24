package controllers

import (
	"mygram/models"
	"mygram/repository"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type commentController struct {
	commentRepository repository.ICommentRepository
}

func NewCommentController(commentRepository repository.ICommentRepository) *commentController {
	return &commentController{
		commentRepository: commentRepository,
	}
}

func (cc *commentController) CreateComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := int(userData["id"].(float64))
	comment := models.Comment{}

	err := ctx.ShouldBindJSON(&comment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Bad request",
			"message": err.Error(),
		})
		return
	}

	comment.UserID = uint(userID)

	createdComment, err := cc.commentRepository.Create(comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, createdComment)
}

func (cc *commentController) GetComments(ctx *gin.Context) {
	comments, err := cc.commentRepository.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (cc *commentController) GetCommentById(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	comment, err := cc.commentRepository.GetOne(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func (cc *commentController) UpdateComment(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	var comment models.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	comment.ID = uint(commentID)

	updatedComment, err := cc.commentRepository.Update(comment, commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, updatedComment)
}

func (cc *commentController) DeleteComment(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentId"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = cc.commentRepository.Delete(commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Delete error",
			"message": err.Error(),
		})
		return
	}
	
	ctx.JSON(http.StatusOK, gin.H{
		"message" : "Your comment has been successfully deleted",
	})
}