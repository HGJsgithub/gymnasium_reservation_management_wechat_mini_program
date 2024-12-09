package tennis

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/common/utils"
	"Gymnasium_reservation_WeChat_mini_program/model/venue"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTennisVenueState1 获取今天的场地状态表
func GetTennisVenueState1(c *gin.Context) {
	db := database.InitGormDB()
	var venueState []venue.TennisVenueState1
	db.Find(&venueState)
	//由于读取出来的数据都是0和1，因此要把0转化为false，1转化为true。rawState就是从数据库读出来的原始数据
	rawState := utils.TennisStructToSlice1(venueState)
	//stateTable是转化完成的状态表
	stateTable := utils.GetVenueStateTable(rawState, 13)
	c.JSON(http.StatusOK, stateTable)
}

// GetTennisVenueState2 获取明天的场地状态表
func GetTennisVenueState2(c *gin.Context) {
	db := database.InitGormDB()
	var venueState []venue.TennisVenueState2
	db.Find(&venueState)
	//由于读取出来的数据都是0和1，因此要把0转化为false，1转化为true。rawState就是从数据库读出来的原始数据
	rawState := utils.TennisStructToSlice2(venueState)
	//stateTable是转化完成的状态表
	stateTable := utils.GetVenueStateTable(rawState, 13)
	c.JSON(http.StatusOK, stateTable)
}
