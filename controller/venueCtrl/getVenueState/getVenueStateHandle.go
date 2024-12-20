package getVenueState

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model/venueModel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBadmintonVenueStateTableHandle(c *gin.Context) {
	db := database.ConnectToGormDB()
	vs := venueModel.BadmintonVenueState{}
	var vsSlice []venueModel.BadmintonVenueState
	db.Where("date = ?", "today").Find(&vsSlice)
	//由于读取出来的数据是结构体数组，且状态都是用0和1来表示，因此要把0转化为false，1转化为true。
	//rawState是经过BadmintonStructToSlice处理后得到的场地状态二维切片数据，但是还包含id和date字段
	rawState := vs.BadmintonStructToSlice(vsSlice)
	fmt.Println("rawState", rawState)
	//vsTable是转化成二维切片的状态表，GetVenueStateSlice用来去除id和date字段
	todayVSTable := RemoveIdAndDate(rawState, 13)
	db.Where("date = ?", "tomorrow").Find(&vsSlice)
	err := db.Close()
	if err != nil {
		fmt.Println("关闭数据库连接错误：", err)
		return
	}
	rawState = vs.BadmintonStructToSlice(vsSlice)
	tomorrowVSTable := RemoveIdAndDate(rawState, 13)
	VSTables := map[string][][]bool{
		"today":    todayVSTable,
		"tomorrow": tomorrowVSTable,
	}
	c.JSON(http.StatusOK, VSTables)
}

func GetTableTennisVenueStateTableHandle(c *gin.Context) {
	db := database.ConnectToGormDB()
	vs := venueModel.TableTennisVenueState{}
	var vsSlice []venueModel.TableTennisVenueState
	db.Where("date = ?", "today").Find(&vsSlice)
	//由于读取出来的数据是结构体数组，且状态都是用0和1来表示，因此要把0转化为false，1转化为true。
	//rawState是经过BadmintonStructToSlice处理后得到的场地状态二维切片数据，但是还包含id和date字段
	rawState := vs.TableTennisStructToSlice(vsSlice)
	fmt.Println("rawState", rawState)
	//vsTable是转化成二维切片的状态表，GetVenueStateSlice用来去除id和date字段
	todayVSTable := RemoveIdAndDate(rawState, 13)
	db.Where("date = ?", "tomorrow").Find(&vsSlice)
	err := db.Close()
	if err != nil {
		fmt.Println("关闭数据库连接错误：", err)
		return
	}
	rawState = vs.TableTennisStructToSlice(vsSlice)
	tomorrowVSTable := RemoveIdAndDate(rawState, 13)
	VSTables := map[string][][]bool{
		"today":    todayVSTable,
		"tomorrow": tomorrowVSTable,
	}
	c.JSON(http.StatusOK, VSTables)
}

func GetTennisVenueStateTableHandle(c *gin.Context) {
	db := database.ConnectToGormDB()
	vs := venueModel.TennisVenueState{}
	var vsSlice []venueModel.TennisVenueState
	db.Where("date = ?", "today").Find(&vsSlice)
	//由于读取出来的数据是结构体数组，且状态都是用0和1来表示，因此要把0转化为false，1转化为true。
	//rawState是经过BadmintonStructToSlice处理后得到的场地状态二维切片数据，但是还包含id和date字段
	rawState := vs.TennisStructToSlice(vsSlice)
	fmt.Println("rawState", rawState)
	//vsTable是转化成二维切片的状态表，GetVenueStateSlice用来去除id和date字段
	todayVSTable := RemoveIdAndDate(rawState, 13)
	db.Where("date = ?", "tomorrow").Find(&vsSlice)
	err := db.Close()
	if err != nil {
		fmt.Println("关闭数据库连接错误：", err)
		return
	}
	rawState = vs.TennisStructToSlice(vsSlice)
	tomorrowVSTable := RemoveIdAndDate(rawState, 13)
	VSTables := map[string][][]bool{
		"today":    todayVSTable,
		"tomorrow": tomorrowVSTable,
	}
	c.JSON(http.StatusOK, VSTables)
}
