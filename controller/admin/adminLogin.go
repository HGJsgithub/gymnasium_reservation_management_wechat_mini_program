package admin

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/common/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminVerification(c *gin.Context) {
	db := database.InitGormDB()
	account := c.PostForm("account")
	password := c.PostForm("password")
	isAccRight := utils.CheckAccount(db, account)
	isPwdRight := utils.CheckAdminPwd(db, password)

	if isAccRight && isPwdRight {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  "账号密码正确",
		})
	}

	if isAccRight == false {
		c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
			"code": 1,
			"msg":  "账号错误！",
		})
	} else if isPwdRight == false {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 2,
			"msg":  "密码错误！",
		})
	}
	defer db.Close()

}
