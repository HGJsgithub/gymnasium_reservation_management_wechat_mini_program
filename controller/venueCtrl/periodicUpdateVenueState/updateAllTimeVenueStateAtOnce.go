package periodicUpdateVenueState

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"strconv"
	"time"
)

func UpdateAllTimeVenueStateAtOnce() {
	nowHour := time.Now().Hour()
	timeField := "t" + strconv.Itoa(nowHour)
	var b venueModel.BadmintonVenueState
	var tt venueModel.TableTennisVenueState
	var t venueModel.TennisVenueState
	if nowHour > 9 {
		for i := 9; i < nowHour; i++ {
			timeField = "t" + strconv.Itoa(i)
			b.UpdateVenueStateEveryHour(timeField)
			tt.UpdateVenueStateEveryHour(timeField)
			t.UpdateVenueStateEveryHour(timeField)
		}
	}
}
