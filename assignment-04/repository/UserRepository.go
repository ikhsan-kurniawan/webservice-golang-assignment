package repository

import (
	"mygram/models"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) Register(newUser models.User) (models.User, error) {
	err := ur.db.Create(&newUser).Error
	return newUser, err
}