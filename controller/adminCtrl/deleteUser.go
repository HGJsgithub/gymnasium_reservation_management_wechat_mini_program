package adminCtrl

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func DeleteUser(c *gin.Context) {
	db := database.ConnectToGormDB()
	phone := c.PostForm("phone")
	user := model.User{
		Phone: phone,
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "成功删除！",
	})
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

}
