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

func (ur *userRepository) Login(user models.User) (models.User, error) {
	var loginUser models.User
	err := ur.db.Where("email = ?", user.Email).Take(&loginUser).Error
	return loginUser, err
}