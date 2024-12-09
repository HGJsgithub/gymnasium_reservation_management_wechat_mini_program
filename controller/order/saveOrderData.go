package order

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func SaveOrderData(c *gin.Context) {
	db := database.ConnectTOGormDB()
	var order model.Order
	err := c.ShouldBindJSON(&order)
	log.Println("这是orderData", order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if db.First(&order).RecordNotFound() == true {
			db.Create(&order)
			c.JSON(http.StatusCreated, gin.H{
				"success": true,
				"msg":     "保存通用订单数据成功！",
				"order":   order,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"msg":     "该订单号已经存在！",
			})
		}
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
}
