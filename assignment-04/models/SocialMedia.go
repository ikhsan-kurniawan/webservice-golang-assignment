package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" valid:"required~name is required"`
	SocialMediaURL string `gorm:"not null" json:"social_media_url" valid:"required~social_media_url is required"`
	UserID         uint		`json:"user_id"`
	User           *User
}

func (socialmedia *SocialMedia) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(socialmedia)

	return err
}

func (socialmedia *SocialMedia) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(socialmedia)

	return err
}