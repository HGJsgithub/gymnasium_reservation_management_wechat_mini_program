package venueStateTable

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"fmt"
	"github.com/jinzhu/gorm"
)

func CreateBadmintonVenueStateData(db *gorm.DB, venueNum int) {
	for i := 0; i < venueNum; i++ {
		todayVenueState := venueModel.BadmintonVenueState{ID: i + 1, Date: "today"}
		tomorrowVenueState := venueModel.BadmintonVenueState{ID: i + 1, Date: "tomorrow"}
		if db.First(&todayVenueState).RecordNotFound() == true {
			db.Create(&venueModel.BadmintonVenueState{
				ID:   i + 1,
				Date: "today",
				T9:   false,
				T10:  false,
				T11:  false,
				T12:  false,
				T13:  false,
				T14:  false,
				T15:  false,
				T16:  false,
				T17:  false,
				T18:  false,
				T19:  false,
				T20:  false,
				T21:  false,
			})
		} else {
			fmt.Println(i+1, "号场地今天的羽毛球场地已经存在！")
			fmt.Println()
		}
		if db.First(&tomorrowVenueState).RecordNotFound() == true {
			db.Create(&venueModel.BadmintonVenueState{
				ID:   i + 1,
				Date: "tomorrow",
				T9:   false,
				T10:  false,
				T11:  false,
				T12:  false,
				T13:  false,
				T14:  false,
				T15:  false,
				T16:  false,
				T17:  false,
				T18:  false,
				T19:  false,
				T20:  false,
				T21:  false,
			})
		} else {
			fmt.Println(i+1, "号场地明天的羽毛球场地已经存在！")
			fmt.Println()
		}
	}
}
