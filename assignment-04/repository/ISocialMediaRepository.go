package repository

import "mygram/models"

type ISocialMediaRepository interface {
	Create(newSocialMedia models.SocialMedia) (models.SocialMedia, error)
	GetAll() ([]models.SocialMedia, error)
	GetOne(socialMediaId int) (models.SocialMedia, error)
	Update(updatedSocialMedia models.SocialMedia, socialMediaId int) (models.SocialMedia, error)
	Delete(socialMediaId int) error
}