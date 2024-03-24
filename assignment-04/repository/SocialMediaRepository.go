package repository

import (
	"mygram/models"

	"gorm.io/gorm"
)

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *socialMediaRepository {
	return &socialMediaRepository{
		db: db,
	}
}

func (smr *socialMediaRepository) Create(newSocialMedia models.SocialMedia) (models.SocialMedia, error) {
	err := smr.db.Create(&newSocialMedia).Error
	return newSocialMedia, err
}

func (smr *socialMediaRepository) GetAll() ([]models.SocialMedia, error) {
	var socialMedias []models.SocialMedia

	if err := smr.db.Preload("User").Find(&socialMedias).Error; err != nil {
		return nil, err
	}
	return socialMedias, nil
}

func (smr *socialMediaRepository) GetOne(socialMediaId int) (models.SocialMedia, error) {
	var socialMedia models.SocialMedia
	if err := smr.db.Preload("User").First(&socialMedia, socialMediaId).Error; err != nil {
		return models.SocialMedia{}, err
	}
	return socialMedia, nil
}

func (smr *socialMediaRepository) Update(updatedSocialMedia models.SocialMedia, socialMediaId int) (models.SocialMedia, error) {
	existedSocialMedia := models.SocialMedia{}
	err := smr.db.First(&existedSocialMedia, socialMediaId).Error
	if err != nil {
		return existedSocialMedia, err
	}

	existedSocialMedia.Name = updatedSocialMedia.Name
	existedSocialMedia.SocialMediaURL = updatedSocialMedia.SocialMediaURL

	err = smr.db.Model(&existedSocialMedia).Where("id = ?", socialMediaId).Updates(existedSocialMedia).Error
	if err != nil {
		return existedSocialMedia, err
	}

	return existedSocialMedia, nil
}

func (smr *socialMediaRepository) Delete(socialMediaId int) error {
	var socialMedia models.SocialMedia

	if err := smr.db.First(&socialMedia, socialMediaId).Error; err != nil {
		return err
	}

	if err := smr.db.Delete(&socialMedia).Error; err != nil {
		return err
	}

	return nil
}
