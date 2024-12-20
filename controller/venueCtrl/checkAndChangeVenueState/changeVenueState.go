package checkAndChangeVenueState

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

func ChangeVenueState(db *gorm.DB, tableName string, todayOrTomorrow string, id int, timeField []string, state bool) {
	for i := range timeField {
		db.Table(tableName).Where("id = ? AND date = ?", id, todayOrTomorrow).Update(timeField[i], state)
		fmt.Println("success")
	}
}
