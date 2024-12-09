package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitGormDB() *gorm.DB {
	driver := "mysql"
	host := "127.0.0.1"
	port := "3306"
	user := "root"
	password := "aaaaaa"
	dbname := "gymnasium_reservation_management_wechat_mini_program"
	charset := "utf8mb4"
	parseTime := "True"
	loc := "Local"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		user, password, host, port, dbname, charset, parseTime, loc)
	db, err := gorm.Open(driver, args)
	if err != nil {
		panic(err)
		fmt.Println("gorm连接mysql出错！")
	}
	return db
}
