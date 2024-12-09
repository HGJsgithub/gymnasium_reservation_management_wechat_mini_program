package admin

import (
	"Gymnasium_reservation_WeChat_mini_program/common/database"
	"Gymnasium_reservation_WeChat_mini_program/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUsers(c *gin.Context) {
	db := database.InitGormDB()
	var userList []model.User
	db.Find(&userList)
	c.JSON(http.StatusOK, userList)
	defer db.Close()

}
