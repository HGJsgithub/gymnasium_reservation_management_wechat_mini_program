package main

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/controller/venueCtrl/periodicUpdateVenueState"
	"Gymnasium_reservation_WeChat_mini_program/routes"
	"Gymnasium_reservation_WeChat_mini_program/venueStateTable"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/robfig/cron/v3"
)

func main() {
	db := database.ConnectToGormDB()

	badmintonVenueNum := 6
	TableTennisVenueNum := 4
	TennisVenueNum := 2
	//创建数据库表并赋值
	venueStateTable.AutoMigrate(db)
	venueStateTable.CreateAnnData(db)
	venueStateTable.CreateBadmintonVenueStateData(db, badmintonVenueNum)
	venueStateTable.CreateTableTennisVenueStateData(db, TableTennisVenueNum)
	venueStateTable.CreateTennisVenueStateData(db, TennisVenueNum)

	//把所有场地状态变成空闲
	//venueStateTable.ResetBadmintonVS(db, 6)
	//venueStateTable.ResetTableTennisVS(db, 4)
	//venueStateTable.ResetTennisVS(db, 2)
	err := db.Close()
	if err != nil {
		fmt.Println("断开数据库连接出错：", err)
	}

	r := gin.Default()
	r = routes.CollectRoute(r)

	c := cron.New()
	periodicUpdateVenueState.UpdateVenueStateEveryHour(c)
	periodicUpdateVenueState.UpdateVenueStateEveryday(c, badmintonVenueNum, TableTennisVenueNum, TennisVenueNum)
	c.Start()

	err = r.Run(":8080")
	if err != nil {
		fmt.Println("r.Run出错：", err)
		return
	}

}
