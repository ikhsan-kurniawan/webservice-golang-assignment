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
	User     *User
	Comment  []Comment `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
}

func (photo *Photo) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(photo)

	return err
}