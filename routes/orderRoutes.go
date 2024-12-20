package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/orderCtrl"
	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.Engine) *gin.Engine {
	//获取订单数据
	r.POST("/order/getOrderData", orderCtrl.GetOrderData)
	//保存订单数据
	r.POST("/order/saveOrderData", orderCtrl.SaveOrderData)
	//改变订单状态
	r.POST("/order/changeOrderState", orderCtrl.ChangeOrderState)
	//改变订单看取消状态
	r.POST("/order/changeCancelFlag", orderCtrl.ChangeCancelFlag)
	return r
}
