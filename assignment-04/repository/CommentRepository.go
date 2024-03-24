package repository

import (
	"mygram/models"

	"gorm.io/gorm"
)

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *commentRepository {
	return &commentRepository{
		db: db,
	}
}

func (cr *commentRepository) Create(newComment models.Comment) (models.Comment, error) {
	var photo models.Photo
    if err := cr.db.First(&photo, newComment.PhotoID).Error; err != nil {
        return models.Comment{}, err
    }

	err := cr.db.Create(&newComment).Error
	return newComment, err
}

func (cr *commentRepository) GetAll() ([]models.Comment, error) {
	var comments []models.Comment

	if err := cr.db.Preload("User").Preload("Photo").Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (cr *commentRepository) GetOne(commentId int) (models.Comment, error) {
	var comment models.Comment
	if err := cr.db.Preload("User").Preload("Photo").First(&comment, commentId).Error; err != nil {
		return models.Comment{}, err
	}
	return comment, nil
}

func (cr *commentRepository) Update(updatedComment models.Comment, commentId int) (models.Comment, error) {
	existedComment := models.Comment{}
	err := cr.db.First(&existedComment, commentId).Error
	if err != nil {
		return existedComment, err
	}

	existedComment.Message = updatedComment.Message

	err = cr.db.Model(&existedComment).Where("id = ?", commentId).Updates(existedComment).Error
	if err != nil {
		return existedComment, err
	}

	return existedComment, nil
}

func (cr *commentRepository) Delete(commentId int) error {
	var comment models.Comment

	if err := cr.db.First(&comment, commentId).Error; err != nil {
		return err
	}

	if err := cr.db.Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}
