package registration

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Registration(c *gin.Context) {
	db := database.ConnectToGormDB()
	var newUser model.User
	err := c.ShouldBindJSON(&newUser)
	fmt.Println(newUser)
	if err != nil {
		c.JSON(http.StatusAccepted, gin.H{"error": err.Error()})
	} else {
		if db.First(&newUser).RecordNotFound() == true {
			db.Create(&newUser)
			c.JSON(http.StatusCreated, gin.H{
				"code":    0,
				"msg":     "注册成功！",
				"newUser": newUser,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "该号码已注册！",
			})
		}
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println("断开数据库连接错误：", err)
		}
	}(db)

}
