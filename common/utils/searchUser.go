package utils

import (
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/jinzhu/gorm"
)

func SearchUser(db *gorm.DB, phone string, password string, userPoint *model.User) (exist, right bool) {
	if IsPhoneExist(db, phone) {
		db.Where("phone = ?", phone).First(userPoint)
		user := *userPoint
		if CheckPwd(password, user.Password) {
			return true, true
		} else {
			return true, false
		}
	} else {
		return false, false
	}

}
