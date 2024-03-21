package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" valid:"required~message is required"`
	UserID  uint
	User    *User
	PhotoID uint
	Photo   *User
}

func (comment *Comment) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(comment)

	return err
}