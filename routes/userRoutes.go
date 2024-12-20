package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/userCtrl/changeNickname"
	"Gymnasium_reservation_WeChat_mini_program/controller/userCtrl/changePassword"
	"Gymnasium_reservation_WeChat_mini_program/controller/userCtrl/registration"
	"Gymnasium_reservation_WeChat_mini_program/controller/userCtrl/userLogin"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	//用户注册
	r.POST("/auth/registration", registration.Registration)
	//用户登录
	r.POST("/auth/login/user", userLogin.UserLogin)
	//修改用户昵称
	r.POST("/changeNickname", changeNickname.ChangeNickname)
	//修改用户密码
	r.POST("/changePassword", changePassword.ChangePassword)
	return r
}
