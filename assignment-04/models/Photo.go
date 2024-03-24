package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	GormModel
	Title    string `gorm:"not null" json:"title" valid:"required~title is required"`
	Caption  string
	PhotoURL string `gorm:"not null" json:"photo_url" valid:"required~photo_url is required"`
	UserID   uint
	User     *User `gorm:"foreignKey:UserID"`
	Comments []Comment `json:"-"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(photo)

	return err
}

func (photo *Photo) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(photo)

	return err
}