package database

import (
	"Gymnasium_reservation_WeChat_mini_program/model"
	"Gymnasium_reservation_WeChat_mini_program/model/venue"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Announcement{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})

	db.AutoMigrate(&venue.BadmintonVenueState1{})
	db.AutoMigrate(&venue.BadmintonVenueState2{})

	db.AutoMigrate(&venue.TableTennisVenueState1{})
	db.AutoMigrate(&venue.TableTennisVenueState2{})

	db.AutoMigrate(&venue.TennisVenueState1{})
	db.AutoMigrate(&venue.TennisVenueState2{})

	db.AutoMigrate(&model.Order{})

}
