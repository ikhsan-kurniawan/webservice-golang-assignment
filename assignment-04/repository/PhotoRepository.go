package repository

import (
	"mygram/models"

	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{
		db: db,
	}
}

func (pr *photoRepository) Create(newPhoto models.Photo) (models.Photo, error) {
	err := pr.db.Create(&newPhoto).Error
	return newPhoto, err
}

func (pr *photoRepository) GetAll() ([]models.Photo, error) {
	var photos []models.Photo

    if err := pr.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age")
	}).Find(&photos).Error; err != nil {
        return nil, err
    }
    return photos, nil
}

func (pr *photoRepository) GetOne(photoId int) (models.Photo, error) {
	var photo models.Photo
    if err := pr.db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id, username, email, age")
	}).First(&photo, photoId).Error; err != nil {
        return models.Photo{}, err
    }
    return photo, nil
}

func (pr *photoRepository) Update(updatedPhoto models.Photo, photoId int) (models.Photo, error) {
	existedPhoto := models.Photo{}
	err := pr.db.First(&existedPhoto, photoId).Error
    if err != nil {
        return existedPhoto, err
    }

	existedPhoto.Title = updatedPhoto.Title
	existedPhoto.Caption = updatedPhoto.Caption
	existedPhoto.PhotoURL = updatedPhoto.PhotoURL

	err = pr.db.Model(&existedPhoto).Where("id = ?", photoId).Updates(existedPhoto).Error
	if err != nil {
		return existedPhoto, err
	}

	return existedPhoto, nil
}

func (pr *photoRepository) Delete(photoId int) error {
	var photo models.Photo

	if err := pr.db.First(&photo, photoId).Error; err != nil {
		return err
	}

	if err := pr.db.Delete(&photo).Error; err != nil {
		return err
	}

	return nil
}
