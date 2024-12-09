package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/announcement"
	"github.com/gin-gonic/gin"
)

func AnnRoute(r *gin.Engine) *gin.Engine {
	//获取公告
	r.GET("/api/announcement/get", announcement.GetAnn)
	//删除公告
	r.POST("/api/announcement/delete", announcement.DeleteAnn)
	return r
}
