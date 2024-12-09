package badminton

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/common/utils"
	"Gymnasium_reservation_WeChat_mini_program/model/venue"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetBadmintonVenueState1 获取今天的场地状态表
func GetBadmintonVenueState1(c *gin.Context) {
	db := database.InitGormDB()
	var venueState []venue.BadmintonVenueState1
	db.Find(&venueState)
	//由于读取出来的数据都是0和1，因此要把0转化为false，1转化为true。rawState就是从数据库读出来的原始数据
	rawState := utils.BadmintonStructToSlice1(venueState)
	//stateTable是转化完成的状态表
	stateTable := utils.GetVenueStateTable(rawState, 13)
	c.JSON(http.StatusOK, stateTable)
}

// GetBadmintonVenueState2 获取明天的场地状态表
func GetBadmintonVenueState2(c *gin.Context) {
	db := database.InitGormDB()
	var venueState []venue.BadmintonVenueState2
	db.Find(&venueState)
	//由于读取出来的数据都是0和1，因此要把0转化为false，1转化为true。rawState就是从数据库读出来的原始数据
	rawState := utils.BadmintonStructToSlice2(venueState)
	//stateTable是转化完成的状态表
	stateTable := utils.GetVenueStateTable(rawState, 13)
	c.JSON(http.StatusOK, stateTable)
}
