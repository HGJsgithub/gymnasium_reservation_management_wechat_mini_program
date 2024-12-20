package checkAndChangeVenueState

import (
	"github.com/jinzhu/gorm"
)

func CheckIfVenueAvailable(db *gorm.DB, tableName string, todayOrTomorrow string, id int, timeField []string) (available bool, msg map[int][]string) {
	//判断场地是否已被预定，例如查询到t9字段为true，即s=true，说明已被预定，并说明哪些场地和时段已被预定
	available = true
	for i := range timeField {
		row, err := db.Table(tableName).Where("id = ? AND date = ?", id, todayOrTomorrow).Select(timeField[i]).Rows()
		if err != nil {
			panic(err)
		}
		var s bool
		for row.Next() {
			err := row.Scan(&s)
			if err != nil {
				panic(err)
			}
			if s == true {
				available = false
				createHasBeenReservedInfo(tableName, id, timeField[i], msg)
			}
		}
	}
	return available, msg
}
