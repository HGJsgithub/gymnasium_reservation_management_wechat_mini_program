package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/admin"
	"github.com/gin-gonic/gin"
)

func AdminRoute(r *gin.Engine) *gin.Engine {
	//管理员登录
	r.POST("/api/auth/login/admin", admin.AdminVerification)
	//获取所有用户
	r.GET("/api/users/get", admin.GetAllUsers)
	//删除用户
	r.POST("/api/users/delete", admin.DeleteUser)
	return r
}
