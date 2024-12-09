package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/user"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) *gin.Engine {
	//用户注册
	r.POST("/api/auth/registration", user.Registration)
	//用户登录
	r.POST("/api/auth/login/user", user.UserLogin)
	//修改用户昵称
	r.POST("/api/changeNickname", user.ChangeNickname)
	//修改用户密码
	r.POST("/api/changePassword", user.ChangePassword)
	return r
}
