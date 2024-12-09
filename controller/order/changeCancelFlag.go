package order

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ChangeCancelFlag(c *gin.Context) {
	db := database.ConnectTOGormDB()
	id := c.PostForm("id")
	state, _ := strconv.ParseBool(c.PostForm("state"))
	if db.Table("orders").Where("id = ?", id).RecordNotFound() == false {
		db.Table("orders").Where("id=?", id).Update("cancel_flag", state)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"msg":     "找到并成功取消状态！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"msg":     "没有找到该订单！",
		})
	}
}
