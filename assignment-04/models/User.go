package models

import (
	"mygram/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username    string        `gorm:"not null;uniqueIndex" json:"username" valid:"required~username is required"`
	Email       string        `gorm:"not null;uniqueIndex" json:"email" valid:"required~email is required,email~Invalid email format"`
	Password    string        `gorm:"not null" json:"password" valid:"required~password is required,minstringlength(6)~password has to have minimum length of 6 characters"`
	Age         int           `gorm:"not null" json:"age" valid:"required~age is required,gt=8~age must greater than 8"`
	Comment     []Comment     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"comments"`
	Photo       []Photo       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photos"`
	SocialMedia []SocialMedia `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"social_medias"`
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	user.Password = helpers.HashPash(user.Password)

	return err

}