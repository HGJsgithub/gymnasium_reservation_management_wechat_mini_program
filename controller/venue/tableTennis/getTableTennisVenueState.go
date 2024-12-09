package tableTennis

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/common/utils"
	"Gymnasium_reservation_WeChat_mini_program/model/venue"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetTableTennisVenueState1 获取今天的场地状态表
func GetTableTennisVenueState1(c *gin.Context) {
	db := database.InitGormDB()
	var venueState []venue.TableTennisVenueState1
	db.Find(&venueState)
	//由于读取出来的数据都是0和1，因此要把0转化为false，1转化为true。rawState就是从数据库读出来的原始数据
	rawState := utils.TableTennisStructToSlice1(venueState)
	//stateTable是转化完成的状态表
	stateTable := utils.GetVenueStateTable(rawState, 13)
	c.JSON(http.StatusOK, stateTable)
}

// GetTableTennisVenueState2 获取明天的场地状态表
func GetTableTennisVenueState2(c *gin.Context) {
	db := database.InitGormDB()
	var venueState []venue.TableTennisVenueState2
	db.Find(&venueState)
	//由于读取出来的数据都是0和1，因此要把0转化为false，1转化为true。rawState就是从数据库读出来的原始数据
	rawState := utils.TableTennisStructToSlice2(venueState)
	//stateTable是转化完成的状态表
	stateTable := utils.GetVenueStateTable(rawState, 13)
	c.JSON(http.StatusOK, stateTable)
}
