package venueStateTable

import (
	"Gymnasium_reservation_WeChat_mini_program/model"
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Announcement{})

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Admin{})

	db.AutoMigrate(&venueModel.BadmintonVenueState{}, &venueModel.TableTennisVenueState{}, &venueModel.TennisVenueState{})

	db.AutoMigrate(&model.Order{})

}
