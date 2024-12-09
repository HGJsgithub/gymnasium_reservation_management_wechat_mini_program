package database

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venue"
	"github.com/jinzhu/gorm"
)

func ResetTVS(db *gorm.DB, venueNum int) {
	for i := 1; i <= venueNum; i++ {
		oneVs1 := venue.TennisVenueState1{ID: i}
		oneVs2 := venue.TennisVenueState2{ID: i}
		db.Save(&oneVs1)
		db.Save(&oneVs2)
	}
}
