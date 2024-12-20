package periodicUpdateVenueState

import (
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"fmt"
	"github.com/robfig/cron/v3"
)

func UpdateVenueStateEveryday(c *cron.Cron, badmintonVenueNum int, tableTennisVenueNum int, tennisVenueNum int) {
	_, err := c.AddFunc("@daily", func() {
		var b venueModel.BadmintonVenueState
		b.UpdateVenueStateEveryday(badmintonVenueNum)
		var tt venueModel.TableTennisVenueState
		tt.UpdateTableTennisVenueStateEveryday(tableTennisVenueNum)
		var t venueModel.TennisVenueState
		t.UpdateTennisVenueStateEveryday(tennisVenueNum)
	})
	if err != nil {
		fmt.Println("每天定期修改场地状态出错：", err)
	}

}
