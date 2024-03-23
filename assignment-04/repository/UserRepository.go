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

func (ur *userRepository) Update(updatedUser models.User, id int) (models.User, error) {
	existedUser := models.User{}
	err := ur.db.First(&existedUser, id).Error
    if err != nil {
        return existedUser, err
    }

	existedUser.Email = updatedUser.Email
	existedUser.Username = updatedUser.Username

	err = ur.db.Model(&existedUser).Where("id = ?", id).Updates(existedUser).Error
	if err != nil {
		return existedUser, err
	}

	return existedUser, nil
}

func (ur *userRepository) Delete(userId int) (error) {
	var user models.User

	if err := ur.db.First(&user, userId).Error; err != nil {
		return err
	}

	if err := ur.db.Delete(&user).Error; err != nil {
		return err
	}

	return nil
}
