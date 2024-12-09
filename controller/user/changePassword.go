package user

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ChangePassword(c *gin.Context) {
	db := database.InitGormDB()
	id := c.PostForm("id")
	password := c.PostForm("password")
	if db.Table("users").Where("id = ?", id).RecordNotFound() == true {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
		})
		return
	}
	db.Table("users").Where("id = ?", id).Update("password", password)
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}
