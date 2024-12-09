package admin

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteUser(c *gin.Context) {
	db := database.InitGormDB()
	phone := c.PostForm("phone")
	user := model.User{
		Phone: phone,
	}
	db.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"msg": "成功删除！",
	})
	defer db.Close()

}
