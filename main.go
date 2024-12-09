package main

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/controller/admin"
	"Gymnasium_reservation_WeChat_mini_program/controller/announcement"
	"Gymnasium_reservation_WeChat_mini_program/controller/order"
	"Gymnasium_reservation_WeChat_mini_program/controller/user"
	"Gymnasium_reservation_WeChat_mini_program/controller/venue/badminton"
	"Gymnasium_reservation_WeChat_mini_program/controller/venue/tableTennis"
	"Gymnasium_reservation_WeChat_mini_program/controller/venue/tennis"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func main() {
	db = database.InitGormDB()
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
	//database.ResetBadVS(db, 10)
	//database.ResetTTVS(db, 8)
	//database.ResetTVS(db, 6)

	r := gin.Default()

	//管理员登录
	r.POST("/api/auth/login/admin", admin.AdminVerification)
	//获取所有用户
	r.GET("/api/users/get", admin.GetAllUsers)
	//删除用户
	r.POST("/api/users/delete", admin.DeleteUser)
	//获取公告
	r.GET("/api/announcement/get", announcement.GetAnn)
	//删除公告
	r.POST("/api/announcement/delete", announcement.DeleteAnn)

	//用户注册
	r.POST("/api/auth/registration", user.Registration)
	//用户登录
	r.POST("/api/auth/login/user", user.UserLogin)
	//修改用户昵称
	r.POST("/api/changeNickname", user.ChangeNickname)
	//修改用户密码
	r.POST("/api/changePassword", user.ChangePassword)

	//羽毛球相关

	//获取今天的场地状态
	r.GET("/api/venueInfo/venueState/badminton/getVenueState/today", badminton.GetBadmintonVenueState1)
	//获取明天的场地状态
	r.GET("/api/venueInfo/venueState/badminton/getVenueState/tomorrow", badminton.GetBadmintonVenueState2)
	//改变今天的场地状态
	r.POST("/api/venueInfo/venueState/badminton/checkAndChangeVenueState/today", badminton.CheckAndChangeBadmintonVenueState1)
	//明天的场地状态
	r.POST("/api/venueInfo/venueState/badminton/checkAndChangeVenueState/tomorrow", badminton.CheckAndChangeBadmintonVenueState2)

	//乒乓球相关
	r.GET("/api/venueInfo/venueState/tableTennis/getVenueState/today", tableTennis.GetTableTennisVenueState1)
	r.GET("/api/venueInfo/venueState/tableTennis/getVenueState/tomorrow", tableTennis.GetTableTennisVenueState2)
	r.POST("/api/venueInfo/venueState/tableTennis/checkAndChangeVenueState/today", tableTennis.CheckAndChangeTableTennisVenueState1)
	r.POST("/api/venueInfo/venueState/tableTennis/checkAndChangeVenueState/tomorrow", tableTennis.CheckAndChangeTableTennisVenueState2)

	//网球相关
	r.GET("/api/venueInfo/venueState/tennis/getVenueState/today", tennis.GetTennisVenueState1)
	r.GET("/api/venueInfo/venueState/tennis/getVenueState/tomorrow", tennis.GetTennisVenueState2)
	r.POST("/api/venueInfo/venueState/tennis/checkAndChangeVenueState/today", tennis.CheckAndChangeTennisVenueState1)
	r.POST("/api/venueInfo/venueState/tennis/checkAndChangeVenueState/tomorrow", tennis.CheckAndChangeTennisVenueState2)

	//订单相关

	//获取订单数据
	r.POST("/api/order/getOrderData", order.GetOrderData)
	//保存订单数据
	r.POST("/api/order/saveOrderData", order.SaveOrderData)
	//改变订单状态
	r.POST("/api/order/changeOrderState", order.ChangeOrderState)
	//改变订单看取消状态
	r.POST("/api/order/changeCancelFlag", order.ChangeCancelFlag)

	err := r.Run(":8080")
	if err != nil {
		return
	}

}
