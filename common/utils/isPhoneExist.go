package utils

import (
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/jinzhu/gorm"
)

func IsPhoneExist(db *gorm.DB, phone string) bool {
	var user model.User
	db.Where("phone = ?", phone).First(&user)
	if user.Phone == phone {
		return true
	} else {
		return false
	}
}
