package venueStateTable

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"github.com/jinzhu/gorm"
)

func ResetBadmintonVS(db *gorm.DB, venueNum int) {
	for i := 1; i <= venueNum; i++ {
		todayVS := venueModel.BadmintonVenueState{ID: i, Date: "today"}
		tomorrowVS := venueModel.BadmintonVenueState{ID: i, Date: "tomorrow"}
		db.Save(&todayVS)
		db.Save(&tomorrowVS)
	}
}
