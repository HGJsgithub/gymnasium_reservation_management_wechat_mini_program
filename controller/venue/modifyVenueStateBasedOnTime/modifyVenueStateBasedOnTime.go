package modifyVenueStateBasedOnTime

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func ModifyVenueStateBasedOnTime(c *cron.Cron) {
	_, err := c.AddFunc("@hourly", func() {
		fmt.Println("Every hour")
	})
	if err != nil {
		return
	}

}
