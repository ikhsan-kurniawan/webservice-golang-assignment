package repository

import "mygram/models"

type ICommentRepository interface {
	Create(newComment models.Comment) (models.Comment, error)
	GetAll() ([]models.Comment, error)
	GetOne(commentId int) (models.Comment, error)
	Update(updatedComment models.Comment, commentId int) (models.Comment, error)
	Delete(commentId int) error
}