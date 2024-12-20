package orderCtrl

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GetOrderData(c *gin.Context) {
	db := database.ConnectToGormDB()
	userID := c.PostForm("userID")
	state := c.PostForm("state")
	log.Println("userID:", userID, "state:", state)
	var orderList []model.Order
	db.Where(map[string]interface{}{"user_id": userID, "state": state}).Find(&orderList)
	if len(orderList) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":    true,
			"msg":       "查询订单数据成功！",
			"orderData": orderList,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"msg":    "什么都没找到！",
		})
	}
}
