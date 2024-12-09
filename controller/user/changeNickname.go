package user

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangeNickname(c *gin.Context) {
	db := database.ConnectTOGormDB()
	id := c.PostForm("id")
	nickname := c.PostForm("nickname")
	if db.Table("users").Where("id = ?", id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
		})
		return
	}
	db.Table("users").Where("id = ?", id).Update("nickname", nickname)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
