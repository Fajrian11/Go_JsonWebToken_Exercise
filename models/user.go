package models

import (
	"Go_JsonWebToken_Exercise/helpers"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	FullName string    `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full name is required"`
	Email    string    `gorm:"not null" json:"email" form:"email" valid:"required~Your Email is required,email~Invalid email format"`
	Password string    `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)"`
	Products []Product `gorm:"constraint:onUpdate:CASCADE,onDelete:SET NULL;" json:"products"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil

	return

}
