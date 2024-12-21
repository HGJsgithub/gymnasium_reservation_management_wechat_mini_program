package periodicUpdateVenueState

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"fmt"
	"github.com/robfig/cron/v3"
	"strconv"
	"time"
)

func UpdateVenueStateEveryHour(c *cron.Cron) {
	_, err := c.AddFunc("@hourly", func() {
		nowHour := time.Now().Hour()
		nowHour--
		timeField := "t" + strconv.Itoa(nowHour)
		var b venueModel.BadmintonVenueState
		var tt venueModel.TableTennisVenueState
		var t venueModel.TennisVenueState
		b.UpdateVenueStateEveryHour(timeField)
		tt.UpdateVenueStateEveryHour(timeField)
		t.UpdateVenueStateEveryHour(timeField)

	})
	if err != nil {
		fmt.Println("每个小时修改场地状态出错：", err)
	}

}
