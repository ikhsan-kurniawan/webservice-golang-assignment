package repository

import "mygram/models"

type IPhotoRepository interface {
	Create(newPhoto models.Photo) (models.Photo, error)
	GetAll() ([]models.Photo, error)
	GetOne(photoId int) (models.Photo, error)
	Update(updatedPhoto models.Photo, photoId int) (models.Photo, error)
	Delete(photoId int) error
}