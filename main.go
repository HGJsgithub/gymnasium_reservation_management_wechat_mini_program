package main

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/controller/venue/modifyVenueStateBasedOnTime"
	"Gymnasium_reservation_WeChat_mini_program/routes"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron/v3"
)

var db *gorm.DB

func main() {
	db = database.ConnectTOGormDB()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	database.AutoMigrate(db)
	database.CreateAnnData(db)
	database.CreateBadmintonVenueStateData(db, 10)
	database.CreateTableTennisVenueStateData(db, 8)
	database.CreateTennisVenueStateData(db, 6)

	//把所有场地状态变成空闲
	database.ResetBadVS(db, 10)
	database.ResetTTVS(db, 8)
	database.ResetTVS(db, 6)

	c := cron.New()

	modifyVenueStateBasedOnTime.ModifyVenueStateBasedOnTime(c)

	c.Start()

	r := gin.Default()
	r = routes.CollectRoute(r)
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"sxy": "你好",
	//	})
	//})
	err := r.Run(":8080")
	if err != nil {
		return
	}

}
