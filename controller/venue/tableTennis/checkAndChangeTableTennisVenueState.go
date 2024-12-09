package tableTennis

import (
	"Gymnasium_reservation_WeChat_mini_program/common/constData"
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func CheckAndChangeTableTennisVenueState1(c *gin.Context) {
	tableName := "table_tennis_venue_state1"
	db := database.ConnectTOGormDB()
	count := c.PostForm("count")
	state, _ := strconv.ParseBool(c.PostForm("state"))
	var success = true
	msg := map[int][]string{}
	strID := c.PostForm("id1")
	id1, _ := strconv.Atoi(strID)
	strTimeField1 := c.PostForm("timeField1")
	timeField1 := strings.Split(strTimeField1, ",")
	//判断是否有场地被提前预约
	for i := range timeField1 {
		row, _ := db.Table(tableName).Where("id=?", id1).Select(timeField1[i]).Rows()
		var s bool
		for row.Next() {
			err := row.Scan(&s)
			if err != nil {
				log.Fatal(err)
			}
			if s == true {
				success = false
			}
		}
		if err := row.Err(); err != nil {
			log.Fatal(err)
		}
	}
	for i := range timeField1 {
		row, _ := db.Table(tableName).Where("id=?", id1).Select(timeField1[i]).Rows()
		var s bool
		for row.Next() {
			err := row.Scan(&s)
			if err != nil {
				log.Fatal(err)
			}
			if s == true {
				fmt.Printf("%d号场地%s时间段已被预订！\n", id1, constData.TableTennisTimePeriod[timeField1[i]])
				msg[id1] = []string{}
				msg[id1] = append(msg[id1], constData.TableTennisTimePeriod[timeField1[i]])
			} else if success == true {
				db.Table(tableName).Where("id = ?", id1).Update(timeField1[i], state)
			}
		}
		if err := row.Err(); err != nil {
			log.Fatal(err)
		}
	}
	if count == "1" {
		if success == false {
			c.JSON(http.StatusOK, gin.H{
				"success": success,
				"msg":     msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": success})
		}
		return
	}
	if count == "2" {
		strID := c.PostForm("id2")
		id2, _ := strconv.Atoi(strID)
		strTimeField2 := c.PostForm("timeField2")
		timeField2 := strings.Split(strTimeField2, ",")
		//判断是否有场地被提前预约
		for i := range timeField2 {
			row, _ := db.Table(tableName).Where("id=?", id2).Select(timeField2[i]).Rows()
			var s bool
			for row.Next() {
				err := row.Scan(&s)
				if err != nil {
					log.Fatal(err)
				}
				if s == true {
					success = false
				}
			}
			if err := row.Err(); err != nil {
				log.Fatal(err)
			}
		}
		for i := range timeField2 {
			//查询对应场地对应时间段状态是否为true并赋值给s
			row, _ := db.Table(tableName).Where("id=?", id2).Select(timeField2[i]).Rows()
			var s bool
			for row.Next() {
				err := row.Scan(&s)
				if err != nil {
					log.Fatal(err)
				}
				if s == true {
					fmt.Printf("%d号场地%s时间段已被预订！\n", id2, timeField2[i])
					msg[id2] = []string{}
					msg[id2] = append(msg[id2], constData.TableTennisTimePeriod[timeField2[i]])
				} else if success == true {
					db.Table(tableName).Where("id = ?", id2).Update(timeField2[i], state)
				}
			}
			if err := row.Err(); err != nil {
				log.Fatal(err)
			}
		}
		if success == false {
			c.JSON(http.StatusOK, gin.H{
				"success": success,
				"msg":     msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": success})
		}
	}
}

func CheckAndChangeTableTennisVenueState2(c *gin.Context) {
	tableName := "table_tennis_venue_state2"
	db := database.ConnectTOGormDB()
	count := c.PostForm("count")
	state, _ := strconv.ParseBool(c.PostForm("state"))
	var success = true
	msg := map[int][]string{}
	strID := c.PostForm("id1")
	id1, _ := strconv.Atoi(strID)
	strTimeField1 := c.PostForm("timeField1")
	timeField1 := strings.Split(strTimeField1, ",")
	//判断是否有场地被提前预约
	for i := range timeField1 {
		row, _ := db.Table(tableName).Where("id=?", id1).Select(timeField1[i]).Rows()
		var s bool
		for row.Next() {
			err := row.Scan(&s)
			if err != nil {
				log.Fatal(err)
			}
			if s == true {
				success = false
			}
		}
		if err := row.Err(); err != nil {
			log.Fatal(err)
		}
	}
	for i := range timeField1 {
		row, _ := db.Table(tableName).Where("id=?", id1).Select(timeField1[i]).Rows()
		var s bool
		for row.Next() {
			err := row.Scan(&s)
			if err != nil {
				log.Fatal(err)
			}
			if s == true {
				fmt.Printf("%d号场地%s时间段已被预订！\n", id1, constData.TableTennisTimePeriod[timeField1[i]])
				msg[id1] = []string{}
				msg[id1] = append(msg[id1], constData.TableTennisTimePeriod[timeField1[i]])
			} else if success == true {
				db.Table(tableName).Where("id = ?", id1).Update(timeField1[i], state)
			}
		}
		if err := row.Err(); err != nil {
			log.Fatal(err)
		}
	}
	if count == "1" {
		if success == false {
			c.JSON(http.StatusOK, gin.H{
				"success": success,
				"msg":     msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": success})
		}
		return
	}
	if count == "2" {
		strID := c.PostForm("id2")
		id2, _ := strconv.Atoi(strID)
		strTimeField2 := c.PostForm("timeField2")
		timeField2 := strings.Split(strTimeField2, ",")
		//判断是否有场地被提前预约
		for i := range timeField2 {
			row, _ := db.Table(tableName).Where("id=?", id2).Select(timeField2[i]).Rows()
			var s bool
			for row.Next() {
				err := row.Scan(&s)
				if err != nil {
					log.Fatal(err)
				}
				if s == true {
					success = false
				}
			}
			if err := row.Err(); err != nil {
				log.Fatal(err)
			}
		}
		for i := range timeField2 {
			//查询对应场地对应时间段状态是否为true并赋值给s
			row, _ := db.Table(tableName).Where("id=?", id2).Select(timeField2[i]).Rows()
			var s bool
			for row.Next() {
				err := row.Scan(&s)
				if err != nil {
					log.Fatal(err)
				}
				if s == true {
					fmt.Printf("%d号场地%s时间段已被预订！\n", id2, timeField2[i])
					msg[id2] = []string{}
					msg[id2] = append(msg[id2], constData.TableTennisTimePeriod[timeField2[i]])
				} else if success == true {
					db.Table(tableName).Where("id = ?", id2).Update(timeField2[i], state)
				}
			}
			if err := row.Err(); err != nil {
				log.Fatal(err)
			}
		}
		if success == false {
			c.JSON(http.StatusOK, gin.H{
				"success": success,
				"msg":     msg,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{"success": success})
		}
	}
}
