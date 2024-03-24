package models

import (
	"errors"
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" valid:"required~username is required"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" valid:"required~email is required,email~Invalid email format"`
	Password    string        `gorm:"not null" json:"password" valid:"required~password is required,minstringlength(6)~password has to have minimum length of 6 characters"`
	Age         int           `gorm:"not null" json:"age" valid:"required~age is required,numeric~age must be numeric"`
	Photos       []Photo       `json:"-"`
	Comments     []Comment     `json:"-"`
	SocialMedias []SocialMedia `json:"-"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	if user.Age <= 8 {
		return errors.New("age must be greater than 8")
	}

	user.Password = helpers.HashPash(user.Password)

	return err

}

func (user *User) BeforeUpdate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(user)

	return err
}