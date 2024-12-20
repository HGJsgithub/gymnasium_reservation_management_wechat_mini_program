package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/annCtrl"
	"github.com/gin-gonic/gin"
)

func AnnRoute(r *gin.Engine) *gin.Engine {
	//获取公告
	r.GET("/announcement/get", annCtrl.GetAnn)
	//删除公告
	r.POST("/announcement/delete", annCtrl.DeleteAnn)
	return r
}
