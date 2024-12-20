package orderCtrl

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeOrderState(c *gin.Context) {
	db := database.ConnectToGormDB()
	id := c.PostForm("id")
	state := c.PostForm("state")
	if db.Table("orders").Where("id = ?", id).RecordNotFound() == false {
		db.Table("orders").Where("id=?", id).Update("state", state)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "找到并成功修改状态！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "修改订单状态时没有找到该订单！",
		})
	}
}
