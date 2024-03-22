package repository

import "mygram/models"

type IUserRepository interface {
	Register(newUser models.User) (models.User, error)
	Login(user models.User) (models.User, error)

	// GetAll() ([]models.User, error)
	// Update(newUser models.User) (models.User, error)
	// Delete(id string) error
}