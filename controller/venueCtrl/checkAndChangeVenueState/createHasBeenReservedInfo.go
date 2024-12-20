package checkAndChangeVenueState

import (
	"Gymnasium_reservation_WeChat_mini_program/common/constData"
	"fmt"
)

func createHasBeenReservedInfo(tableName string, id int, time string, msg map[int][]string) {
	var TimePeriod string
	switch tableName {
	case "badminton_venue_states":
		TimePeriod = constData.BadmintonTimePeriod[time]
	case "table_tennis_venue_states":
		TimePeriod = constData.TableTennisTimePeriod[time]
	case "tennis_venue_states":
		TimePeriod = constData.TennisTimePeriod[time]
	}
	fmt.Printf("%d号场地%s时间段已被预订！\n", id, TimePeriod)
	msg[id] = []string{}
	msg[id] = append(msg[id], TimePeriod)
}
