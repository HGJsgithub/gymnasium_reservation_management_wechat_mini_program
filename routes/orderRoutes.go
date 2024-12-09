package routes

import (
	"Gymnasium_reservation_WeChat_mini_program/controller/order"
	"github.com/gin-gonic/gin"
)

func OrderRoute(r *gin.Engine) *gin.Engine {
	//获取订单数据
	r.POST("/api/order/getOrderData", order.GetOrderData)
	//保存订单数据
	r.POST("/api/order/saveOrderData", order.SaveOrderData)
	//改变订单状态
	r.POST("/api/order/changeOrderState", order.ChangeOrderState)
	//改变订单看取消状态
	r.POST("/api/order/changeCancelFlag", order.ChangeCancelFlag)
	return r
}
