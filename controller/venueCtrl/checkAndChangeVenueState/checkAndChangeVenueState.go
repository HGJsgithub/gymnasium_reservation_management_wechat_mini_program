package checkAndChangeVenueState

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/controller/venueCtrl/venueTypeToTableName"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
)

func CheckAndChangeVenueState(c *gin.Context) {
	db := database.ConnectToGormDB()

	venueType := c.PostForm("venueType")
	fmt.Println("venueType", venueType)
	tableName := venueTypeToTableName.VenueTypeToTableName(venueType)
	todayOrTomorrow := c.PostForm("todayOrTomorrow")
	fmt.Println("todayOrTomorrow", todayOrTomorrow)
	count := c.PostForm("count")
	fmt.Println("count", count)

	state, _ := strconv.ParseBool(c.PostForm("state"))
	fmt.Println("state", state)

	strID := c.PostForm("id1")

	id1, _ := strconv.Atoi(strID)
	fmt.Println("id1", id1)

	strTimeField1 := c.PostForm("timeField1")
	//timeField1例：["t9","t10"]
	timeField1 := strings.Split(strTimeField1, ",")
	fmt.Println("timeField1", timeField1)

	available, msg := CheckIfVenueAvailable(db, tableName, todayOrTomorrow, id1, timeField1)
	fmt.Println("available", available)
	fmt.Println("msg", msg)

	if available {
		ChangeVenueState(db, tableName, todayOrTomorrow, id1, timeField1, state)
	}
	if count == "1" {
		if available == false {
			c.JSON(http.StatusOK, gin.H{
				"success":       available,
				"false message": msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": available})
		}
		return
	}
	if count == "2" {
		strID = c.PostForm("id2")
		id2, _ := strconv.Atoi(strID)
		strTimeField2 := c.PostForm("timeField2")
		timeField2 := strings.Split(strTimeField2, ",")

		available, msg = CheckIfVenueAvailable(db, tableName, todayOrTomorrow, id2, timeField2)

		ChangeVenueState(db, tableName, todayOrTomorrow, id2, timeField2, state)

		if available == false {
			c.JSON(http.StatusOK, gin.H{
				"success": available,
				"msg":     msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": available})
		}
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("断开数据库连接错误：", err)
		}
	}(db)
}
