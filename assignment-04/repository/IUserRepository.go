package repository

import "mygram/models"

type IUserRepository interface {
	Register(newUser models.User) (models.User, error)
	Login(user models.User) (models.User, error)
	Update(updatedUser models.User, id int) (models.User, error)
	Delete(userId int) (error)
}