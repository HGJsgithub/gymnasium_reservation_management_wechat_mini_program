package user

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/common/utils"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLogin(c *gin.Context) {
	db := database.ConnectTOGormDB()
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	println("phone:", phone, "password:", password)
	var user model.User
	isUserExist, isPasswordRight := utils.SearchUser(db, phone, password, &user)
	if isUserExist == true && isPasswordRight == true {
		c.JSON(http.StatusOK, gin.H{
			"code":     0,
			"msg":      "登录成功！",
			"userInfo": user,
		})
	} else if isUserExist == true && isPasswordRight == false {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "密码错误！",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2,
			"msg":  "您还没有注册！",
		})
	}
}
