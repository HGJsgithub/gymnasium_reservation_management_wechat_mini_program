package adminLogin

import (
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/jinzhu/gorm"
)

func CheckAccount(db *gorm.DB, account string) bool {
	var admin model.Admin
	db.Where("account = ?", account).First(&admin)
	if admin.Account == account {
		return true
	} else {
		return false
	}

}

func CheckAdminPwd(db *gorm.DB, password string) bool {
	var admin model.Admin
	db.Where("password = ?", password).First(&admin)
	if admin.Password == password {
		return true
	} else {
		return false
	}
}

//func checkAdmin(db *gorm.DB, account string, password string, adminPoint *model.Admin) (accRight, pwdRight bool) {
//	if CheckAccount(db, account) {
//		db.Where("account = ?", account).First(adminPoint)
//		adminCtrl := *adminPoint
//		if CheckPwd(password, adminCtrl.Password) {
//			return true, true
//		} else {
//			return true, false
//		}
//	} else {
//		return false, false
//	}
//
//}
